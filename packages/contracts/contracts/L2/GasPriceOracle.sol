// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Predeploys } from "../libraries/Predeploys.sol";
import { L1Block } from "../L2/L1Block.sol";
import { ISemver } from "../universal/ISemver.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x4200000000000000000000000000000000000005
 * @title GasPriceOracle
 * @notice This contract maintains the variables responsible for computing the L1 portion of the
 *         total fee charged on L2. It also exposes an API that is useful for knowing how large
 *         the L1 portion of the transaction fee will be.
 */
contract GasPriceOracle is ISemver {
    /**
     * @notice Number of decimals used in the scalar.
     */
    uint256 public constant DECIMALS = 6;

    /**
     * @notice Semantic version.
     * @custom:semver 1.0.0
     */
    string public constant version = "1.0.0";

    /**
     * @notice Computes the L1 portion of the fee based on the size of the rlp encoded input
     *         transaction, the current L1 base fee, and the various dynamic parameters.
     *
     * @param _data Unsigned fully RLP-encoded transaction to get the L1 fee for.
     *
     * @return L1 fee that should be paid for the tx
     */
    function getL1Fee(bytes memory _data) external view returns (uint256) {
        uint256 l1GasUsed = getL1GasUsed(_data);
        uint256 l1Fee = l1GasUsed * l1BaseFee();
        uint256 divisor = 10**DECIMALS;
        uint256 unscaled = l1Fee * scalar();
        uint256 scaled = unscaled / divisor;
        return scaled;
    }

    /**
     * @notice Retrieves the current gas price (base fee).
     *
     * @return Current L2 gas price (base fee).
     */
    function gasPrice() public view returns (uint256) {
        return block.basefee;
    }

    /**
     * @notice Retrieves the current base fee.
     *
     * @return Current L2 base fee.
     */
    function baseFee() public view returns (uint256) {
        return block.basefee;
    }

    /**
     * @notice Retrieves the current fee overhead.
     *
     * @return Current fee overhead.
     */
    function overhead() public view returns (uint256) {
        return L1Block(Predeploys.L1_BLOCK_ATTRIBUTES).l1FeeOverhead();
    }

    /**
     * @notice Retrieves the current fee scalar.
     *
     * @return Current fee scalar.
     */
    function scalar() public view returns (uint256) {
        return L1Block(Predeploys.L1_BLOCK_ATTRIBUTES).l1FeeScalar();
    }

    /**
     * @notice Retrieves the latest known L1 base fee.
     *
     * @return Latest known L1 base fee.
     */
    function l1BaseFee() public view returns (uint256) {
        return L1Block(Predeploys.L1_BLOCK_ATTRIBUTES).basefee();
    }

    /**
     * @notice Computes the amount of L1 gas used for a transaction. Adds the overhead which
     *         represents the per-transaction gas overhead of posting the transaction and state
     *         roots to L1. Adds 68 bytes of padding to account for the fact that the input does
     *         not have a signature.
     *
     * @param _data Unsigned fully RLP-encoded transaction to get the L1 gas for.
     *
     * @return Amount of L1 gas used to publish the transaction.
     */
    function getL1GasUsed(bytes memory _data) public view returns (uint256) {
        uint256 total = 0;
        uint256 length = _data.length;
        for (uint256 i = 0; i < length; i++) {
            if (_data[i] == 0) {
                total += 4;
            } else {
                total += 16;
            }
        }
        return total + overhead();
    }
}
