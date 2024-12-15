// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MarketAllocation is an auto generated low-level Go binding around an user-defined struct.
type MarketAllocation struct {
	MarketParams MarketParams
	Assets       *big.Int
}

// MarketParams is an auto generated low-level Go binding around an user-defined struct.
type MarketParams struct {
	LoanToken       common.Address
	CollateralToken common.Address
	Oracle          common.Address
	Irm             common.Address
	Lltv            *big.Int
}

// MetaMorphoVaultMetaData contains all meta data concerning the MetaMorphoVault contract.
var MetaMorphoVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"morpho\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialTimelock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AboveMaxTimelock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllCapsReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BelowMinTimelock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"DuplicateMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InconsistentAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InconsistentReallocation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InvalidMarketRemovalNonZeroCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InvalidMarketRemovalNonZeroSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InvalidMarketRemovalTimelockNotElapsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketNotCreated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MarketNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MathOverflowedMulDiv\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxQueueLengthExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPendingValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllocatorRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCuratorNorGuardianRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCuratorRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotGuardianRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"PendingCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRemoval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"SupplyCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotElapsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"UnauthorizedMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFeeRecipient\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeShares\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"suppliedAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"suppliedShares\",\"type\":\"uint256\"}],\"name\":\"ReallocateSupply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnShares\",\"type\":\"uint256\"}],\"name\":\"ReallocateWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RevokePendingCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RevokePendingGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RevokePendingMarketRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RevokePendingTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"SetCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCurator\",\"type\":\"address\"}],\"name\":\"SetCurator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"SetFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"SetGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllocator\",\"type\":\"bool\"}],\"name\":\"SetIsAllocator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSkimRecipient\",\"type\":\"address\"}],\"name\":\"SetSkimRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"Id[]\",\"name\":\"newSupplyQueue\",\"type\":\"bytes32[]\"}],\"name\":\"SetSupplyQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"SetTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"Id[]\",\"name\":\"newWithdrawQueue\",\"type\":\"bytes32[]\"}],\"name\":\"SetWithdrawQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Skim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"SubmitCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"SubmitGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"SubmitMarketRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"SubmitTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedTotalAssets\",\"type\":\"uint256\"}],\"name\":\"UpdateLastTotalAssets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DECIMALS_OFFSET\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MORPHO\",\"outputs\":[{\"internalType\":\"contractIMorpho\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"}],\"name\":\"acceptCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"config\",\"outputs\":[{\"internalType\":\"uint184\",\"name\":\"cap\",\"type\":\"uint184\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"removableAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAllocator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTotalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingCap\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"value\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingTimelock\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"value\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"internalType\":\"structMarketAllocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"name\":\"reallocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"revokePendingCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"revokePendingMarketRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCurator\",\"type\":\"address\"}],\"name\":\"setCurator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAllocator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"newIsAllocator\",\"type\":\"bool\"}],\"name\":\"setIsAllocator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSkimRecipient\",\"type\":\"address\"}],\"name\":\"setSkimRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id[]\",\"name\":\"newSupplyQueue\",\"type\":\"bytes32[]\"}],\"name\":\"setSupplyQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"skimRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newSupplyCap\",\"type\":\"uint256\"}],\"name\":\"submitCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"submitGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"}],\"name\":\"submitMarketRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"submitTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"supplyQueue\",\"outputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"updateWithdrawQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawQueue\",\"outputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MetaMorphoVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use MetaMorphoVaultMetaData.ABI instead.
var MetaMorphoVaultABI = MetaMorphoVaultMetaData.ABI

// MetaMorphoVault is an auto generated Go binding around an Ethereum contract.
type MetaMorphoVault struct {
	MetaMorphoVaultCaller     // Read-only binding to the contract
	MetaMorphoVaultTransactor // Write-only binding to the contract
	MetaMorphoVaultFilterer   // Log filterer for contract events
}

// MetaMorphoVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaMorphoVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaMorphoVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaMorphoVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaMorphoVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetaMorphoVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaMorphoVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaMorphoVaultSession struct {
	Contract     *MetaMorphoVault  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaMorphoVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaMorphoVaultCallerSession struct {
	Contract *MetaMorphoVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MetaMorphoVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaMorphoVaultTransactorSession struct {
	Contract     *MetaMorphoVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MetaMorphoVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaMorphoVaultRaw struct {
	Contract *MetaMorphoVault // Generic contract binding to access the raw methods on
}

// MetaMorphoVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaMorphoVaultCallerRaw struct {
	Contract *MetaMorphoVaultCaller // Generic read-only contract binding to access the raw methods on
}

// MetaMorphoVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaMorphoVaultTransactorRaw struct {
	Contract *MetaMorphoVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaMorphoVault creates a new instance of MetaMorphoVault, bound to a specific deployed contract.
func NewMetaMorphoVault(address common.Address, backend bind.ContractBackend) (*MetaMorphoVault, error) {
	contract, err := bindMetaMorphoVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVault{MetaMorphoVaultCaller: MetaMorphoVaultCaller{contract: contract}, MetaMorphoVaultTransactor: MetaMorphoVaultTransactor{contract: contract}, MetaMorphoVaultFilterer: MetaMorphoVaultFilterer{contract: contract}}, nil
}

// NewMetaMorphoVaultCaller creates a new read-only instance of MetaMorphoVault, bound to a specific deployed contract.
func NewMetaMorphoVaultCaller(address common.Address, caller bind.ContractCaller) (*MetaMorphoVaultCaller, error) {
	contract, err := bindMetaMorphoVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultCaller{contract: contract}, nil
}

// NewMetaMorphoVaultTransactor creates a new write-only instance of MetaMorphoVault, bound to a specific deployed contract.
func NewMetaMorphoVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaMorphoVaultTransactor, error) {
	contract, err := bindMetaMorphoVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultTransactor{contract: contract}, nil
}

// NewMetaMorphoVaultFilterer creates a new log filterer instance of MetaMorphoVault, bound to a specific deployed contract.
func NewMetaMorphoVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*MetaMorphoVaultFilterer, error) {
	contract, err := bindMetaMorphoVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultFilterer{contract: contract}, nil
}

// bindMetaMorphoVault binds a generic wrapper to an already deployed contract.
func bindMetaMorphoVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MetaMorphoVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaMorphoVault *MetaMorphoVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaMorphoVault.Contract.MetaMorphoVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaMorphoVault *MetaMorphoVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.MetaMorphoVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaMorphoVault *MetaMorphoVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.MetaMorphoVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaMorphoVault *MetaMorphoVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaMorphoVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaMorphoVault *MetaMorphoVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaMorphoVault *MetaMorphoVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.contract.Transact(opts, method, params...)
}

// DECIMALSOFFSET is a free data retrieval call binding the contract method 0xaea70acc.
//
// Solidity: function DECIMALS_OFFSET() view returns(uint8)
func (_MetaMorphoVault *MetaMorphoVaultCaller) DECIMALSOFFSET(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "DECIMALS_OFFSET")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DECIMALSOFFSET is a free data retrieval call binding the contract method 0xaea70acc.
//
// Solidity: function DECIMALS_OFFSET() view returns(uint8)
func (_MetaMorphoVault *MetaMorphoVaultSession) DECIMALSOFFSET() (uint8, error) {
	return _MetaMorphoVault.Contract.DECIMALSOFFSET(&_MetaMorphoVault.CallOpts)
}

// DECIMALSOFFSET is a free data retrieval call binding the contract method 0xaea70acc.
//
// Solidity: function DECIMALS_OFFSET() view returns(uint8)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) DECIMALSOFFSET() (uint8, error) {
	return _MetaMorphoVault.Contract.DECIMALSOFFSET(&_MetaMorphoVault.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MetaMorphoVault *MetaMorphoVaultCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MetaMorphoVault *MetaMorphoVaultSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MetaMorphoVault.Contract.DOMAINSEPARATOR(&_MetaMorphoVault.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MetaMorphoVault.Contract.DOMAINSEPARATOR(&_MetaMorphoVault.CallOpts)
}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCaller) MORPHO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "MORPHO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultSession) MORPHO() (common.Address, error) {
	return _MetaMorphoVault.Contract.MORPHO(&_MetaMorphoVault.CallOpts)
}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) MORPHO() (common.Address, error) {
	return _MetaMorphoVault.Contract.MORPHO(&_MetaMorphoVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MetaMorphoVault.Contract.Allowance(&_MetaMorphoVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MetaMorphoVault.Contract.Allowance(&_MetaMorphoVault.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultSession) Asset() (common.Address, error) {
	return _MetaMorphoVault.Contract.Asset(&_MetaMorphoVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) Asset() (common.Address, error) {
	return _MetaMorphoVault.Contract.Asset(&_MetaMorphoVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MetaMorphoVault.Contract.BalanceOf(&_MetaMorphoVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MetaMorphoVault.Contract.BalanceOf(&_MetaMorphoVault.CallOpts, account)
}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) view returns(uint184 cap, bool enabled, uint64 removableAt)
func (_MetaMorphoVault *MetaMorphoVaultCaller) Config(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Cap         *big.Int
	Enabled     bool
	RemovableAt uint64
}, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "config", arg0)

	outstruct := new(struct {
		Cap         *big.Int
		Enabled     bool
		RemovableAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cap = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Enabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RemovableAt = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) view returns(uint184 cap, bool enabled, uint64 removableAt)
func (_MetaMorphoVault *MetaMorphoVaultSession) Config(arg0 [32]byte) (struct {
	Cap         *big.Int
	Enabled     bool
	RemovableAt uint64
}, error) {
	return _MetaMorphoVault.Contract.Config(&_MetaMorphoVault.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) view returns(uint184 cap, bool enabled, uint64 removableAt)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) Config(arg0 [32]byte) (struct {
	Cap         *big.Int
	Enabled     bool
	RemovableAt uint64
}, error) {
	return _MetaMorphoVault.Contract.Config(&_MetaMorphoVault.CallOpts, arg0)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _MetaMorphoVault.Contract.ConvertToAssets(&_MetaMorphoVault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _MetaMorphoVault.Contract.ConvertToAssets(&_MetaMorphoVault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _MetaMorphoVault.Contract.ConvertToShares(&_MetaMorphoVault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _MetaMorphoVault.Contract.ConvertToShares(&_MetaMorphoVault.CallOpts, assets)
}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCaller) Curator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "curator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultSession) Curator() (common.Address, error) {
	return _MetaMorphoVault.Contract.Curator(&_MetaMorphoVault.CallOpts)
}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) Curator() (common.Address, error) {
	return _MetaMorphoVault.Contract.Curator(&_MetaMorphoVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MetaMorphoVault *MetaMorphoVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MetaMorphoVault *MetaMorphoVaultSession) Decimals() (uint8, error) {
	return _MetaMorphoVault.Contract.Decimals(&_MetaMorphoVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) Decimals() (uint8, error) {
	return _MetaMorphoVault.Contract.Decimals(&_MetaMorphoVault.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MetaMorphoVault *MetaMorphoVaultCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MetaMorphoVault *MetaMorphoVaultSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _MetaMorphoVault.Contract.Eip712Domain(&_MetaMorphoVault.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _MetaMorphoVault.Contract.Eip712Domain(&_MetaMorphoVault.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint96)
func (_MetaMorphoVault *MetaMorphoVaultCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint96)
func (_MetaMorphoVault *MetaMorphoVaultSession) Fee() (*big.Int, error) {
	return _MetaMorphoVault.Contract.Fee(&_MetaMorphoVault.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint96)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) Fee() (*big.Int, error) {
	return _MetaMorphoVault.Contract.Fee(&_MetaMorphoVault.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultSession) FeeRecipient() (common.Address, error) {
	return _MetaMorphoVault.Contract.FeeRecipient(&_MetaMorphoVault.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) FeeRecipient() (common.Address, error) {
	return _MetaMorphoVault.Contract.FeeRecipient(&_MetaMorphoVault.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCaller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultSession) Guardian() (common.Address, error) {
	return _MetaMorphoVault.Contract.Guardian(&_MetaMorphoVault.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) Guardian() (common.Address, error) {
	return _MetaMorphoVault.Contract.Guardian(&_MetaMorphoVault.CallOpts)
}

// IsAllocator is a free data retrieval call binding the contract method 0x4dedf20e.
//
// Solidity: function isAllocator(address ) view returns(bool)
func (_MetaMorphoVault *MetaMorphoVaultCaller) IsAllocator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "isAllocator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllocator is a free data retrieval call binding the contract method 0x4dedf20e.
//
// Solidity: function isAllocator(address ) view returns(bool)
func (_MetaMorphoVault *MetaMorphoVaultSession) IsAllocator(arg0 common.Address) (bool, error) {
	return _MetaMorphoVault.Contract.IsAllocator(&_MetaMorphoVault.CallOpts, arg0)
}

// IsAllocator is a free data retrieval call binding the contract method 0x4dedf20e.
//
// Solidity: function isAllocator(address ) view returns(bool)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) IsAllocator(arg0 common.Address) (bool, error) {
	return _MetaMorphoVault.Contract.IsAllocator(&_MetaMorphoVault.CallOpts, arg0)
}

// LastTotalAssets is a free data retrieval call binding the contract method 0x568efc07.
//
// Solidity: function lastTotalAssets() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) LastTotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "lastTotalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTotalAssets is a free data retrieval call binding the contract method 0x568efc07.
//
// Solidity: function lastTotalAssets() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) LastTotalAssets() (*big.Int, error) {
	return _MetaMorphoVault.Contract.LastTotalAssets(&_MetaMorphoVault.CallOpts)
}

// LastTotalAssets is a free data retrieval call binding the contract method 0x568efc07.
//
// Solidity: function lastTotalAssets() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) LastTotalAssets() (*big.Int, error) {
	return _MetaMorphoVault.Contract.LastTotalAssets(&_MetaMorphoVault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _MetaMorphoVault.Contract.MaxDeposit(&_MetaMorphoVault.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _MetaMorphoVault.Contract.MaxDeposit(&_MetaMorphoVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _MetaMorphoVault.Contract.MaxMint(&_MetaMorphoVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _MetaMorphoVault.Contract.MaxMint(&_MetaMorphoVault.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _MetaMorphoVault.Contract.MaxRedeem(&_MetaMorphoVault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _MetaMorphoVault.Contract.MaxRedeem(&_MetaMorphoVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 assets)
func (_MetaMorphoVault *MetaMorphoVaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 assets)
func (_MetaMorphoVault *MetaMorphoVaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _MetaMorphoVault.Contract.MaxWithdraw(&_MetaMorphoVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 assets)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _MetaMorphoVault.Contract.MaxWithdraw(&_MetaMorphoVault.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MetaMorphoVault *MetaMorphoVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MetaMorphoVault *MetaMorphoVaultSession) Name() (string, error) {
	return _MetaMorphoVault.Contract.Name(&_MetaMorphoVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) Name() (string, error) {
	return _MetaMorphoVault.Contract.Name(&_MetaMorphoVault.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) Nonces(owner common.Address) (*big.Int, error) {
	return _MetaMorphoVault.Contract.Nonces(&_MetaMorphoVault.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _MetaMorphoVault.Contract.Nonces(&_MetaMorphoVault.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultSession) Owner() (common.Address, error) {
	return _MetaMorphoVault.Contract.Owner(&_MetaMorphoVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) Owner() (common.Address, error) {
	return _MetaMorphoVault.Contract.Owner(&_MetaMorphoVault.CallOpts)
}

// PendingCap is a free data retrieval call binding the contract method 0xa31be5d6.
//
// Solidity: function pendingCap(bytes32 ) view returns(uint192 value, uint64 validAt)
func (_MetaMorphoVault *MetaMorphoVaultCaller) PendingCap(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "pendingCap", arg0)

	outstruct := new(struct {
		Value   *big.Int
		ValidAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// PendingCap is a free data retrieval call binding the contract method 0xa31be5d6.
//
// Solidity: function pendingCap(bytes32 ) view returns(uint192 value, uint64 validAt)
func (_MetaMorphoVault *MetaMorphoVaultSession) PendingCap(arg0 [32]byte) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _MetaMorphoVault.Contract.PendingCap(&_MetaMorphoVault.CallOpts, arg0)
}

// PendingCap is a free data retrieval call binding the contract method 0xa31be5d6.
//
// Solidity: function pendingCap(bytes32 ) view returns(uint192 value, uint64 validAt)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) PendingCap(arg0 [32]byte) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _MetaMorphoVault.Contract.PendingCap(&_MetaMorphoVault.CallOpts, arg0)
}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns(address value, uint64 validAt)
func (_MetaMorphoVault *MetaMorphoVaultCaller) PendingGuardian(opts *bind.CallOpts) (struct {
	Value   common.Address
	ValidAt uint64
}, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "pendingGuardian")

	outstruct := new(struct {
		Value   common.Address
		ValidAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ValidAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns(address value, uint64 validAt)
func (_MetaMorphoVault *MetaMorphoVaultSession) PendingGuardian() (struct {
	Value   common.Address
	ValidAt uint64
}, error) {
	return _MetaMorphoVault.Contract.PendingGuardian(&_MetaMorphoVault.CallOpts)
}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns(address value, uint64 validAt)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) PendingGuardian() (struct {
	Value   common.Address
	ValidAt uint64
}, error) {
	return _MetaMorphoVault.Contract.PendingGuardian(&_MetaMorphoVault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultSession) PendingOwner() (common.Address, error) {
	return _MetaMorphoVault.Contract.PendingOwner(&_MetaMorphoVault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) PendingOwner() (common.Address, error) {
	return _MetaMorphoVault.Contract.PendingOwner(&_MetaMorphoVault.CallOpts)
}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns(uint192 value, uint64 validAt)
func (_MetaMorphoVault *MetaMorphoVaultCaller) PendingTimelock(opts *bind.CallOpts) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "pendingTimelock")

	outstruct := new(struct {
		Value   *big.Int
		ValidAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns(uint192 value, uint64 validAt)
func (_MetaMorphoVault *MetaMorphoVaultSession) PendingTimelock() (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _MetaMorphoVault.Contract.PendingTimelock(&_MetaMorphoVault.CallOpts)
}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns(uint192 value, uint64 validAt)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) PendingTimelock() (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _MetaMorphoVault.Contract.PendingTimelock(&_MetaMorphoVault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _MetaMorphoVault.Contract.PreviewDeposit(&_MetaMorphoVault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _MetaMorphoVault.Contract.PreviewDeposit(&_MetaMorphoVault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _MetaMorphoVault.Contract.PreviewMint(&_MetaMorphoVault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _MetaMorphoVault.Contract.PreviewMint(&_MetaMorphoVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _MetaMorphoVault.Contract.PreviewRedeem(&_MetaMorphoVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _MetaMorphoVault.Contract.PreviewRedeem(&_MetaMorphoVault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _MetaMorphoVault.Contract.PreviewWithdraw(&_MetaMorphoVault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _MetaMorphoVault.Contract.PreviewWithdraw(&_MetaMorphoVault.CallOpts, assets)
}

// SkimRecipient is a free data retrieval call binding the contract method 0x388af5b5.
//
// Solidity: function skimRecipient() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCaller) SkimRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "skimRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SkimRecipient is a free data retrieval call binding the contract method 0x388af5b5.
//
// Solidity: function skimRecipient() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultSession) SkimRecipient() (common.Address, error) {
	return _MetaMorphoVault.Contract.SkimRecipient(&_MetaMorphoVault.CallOpts)
}

// SkimRecipient is a free data retrieval call binding the contract method 0x388af5b5.
//
// Solidity: function skimRecipient() view returns(address)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) SkimRecipient() (common.Address, error) {
	return _MetaMorphoVault.Contract.SkimRecipient(&_MetaMorphoVault.CallOpts)
}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(bytes32)
func (_MetaMorphoVault *MetaMorphoVaultCaller) SupplyQueue(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "supplyQueue", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(bytes32)
func (_MetaMorphoVault *MetaMorphoVaultSession) SupplyQueue(arg0 *big.Int) ([32]byte, error) {
	return _MetaMorphoVault.Contract.SupplyQueue(&_MetaMorphoVault.CallOpts, arg0)
}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(bytes32)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) SupplyQueue(arg0 *big.Int) ([32]byte, error) {
	return _MetaMorphoVault.Contract.SupplyQueue(&_MetaMorphoVault.CallOpts, arg0)
}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) SupplyQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "supplyQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) SupplyQueueLength() (*big.Int, error) {
	return _MetaMorphoVault.Contract.SupplyQueueLength(&_MetaMorphoVault.CallOpts)
}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) SupplyQueueLength() (*big.Int, error) {
	return _MetaMorphoVault.Contract.SupplyQueueLength(&_MetaMorphoVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MetaMorphoVault *MetaMorphoVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MetaMorphoVault *MetaMorphoVaultSession) Symbol() (string, error) {
	return _MetaMorphoVault.Contract.Symbol(&_MetaMorphoVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) Symbol() (string, error) {
	return _MetaMorphoVault.Contract.Symbol(&_MetaMorphoVault.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) Timelock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) Timelock() (*big.Int, error) {
	return _MetaMorphoVault.Contract.Timelock(&_MetaMorphoVault.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) Timelock() (*big.Int, error) {
	return _MetaMorphoVault.Contract.Timelock(&_MetaMorphoVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 assets)
func (_MetaMorphoVault *MetaMorphoVaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 assets)
func (_MetaMorphoVault *MetaMorphoVaultSession) TotalAssets() (*big.Int, error) {
	return _MetaMorphoVault.Contract.TotalAssets(&_MetaMorphoVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 assets)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) TotalAssets() (*big.Int, error) {
	return _MetaMorphoVault.Contract.TotalAssets(&_MetaMorphoVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) TotalSupply() (*big.Int, error) {
	return _MetaMorphoVault.Contract.TotalSupply(&_MetaMorphoVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _MetaMorphoVault.Contract.TotalSupply(&_MetaMorphoVault.CallOpts)
}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(bytes32)
func (_MetaMorphoVault *MetaMorphoVaultCaller) WithdrawQueue(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "withdrawQueue", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(bytes32)
func (_MetaMorphoVault *MetaMorphoVaultSession) WithdrawQueue(arg0 *big.Int) ([32]byte, error) {
	return _MetaMorphoVault.Contract.WithdrawQueue(&_MetaMorphoVault.CallOpts, arg0)
}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(bytes32)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) WithdrawQueue(arg0 *big.Int) ([32]byte, error) {
	return _MetaMorphoVault.Contract.WithdrawQueue(&_MetaMorphoVault.CallOpts, arg0)
}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCaller) WithdrawQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaMorphoVault.contract.Call(opts, &out, "withdrawQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultSession) WithdrawQueueLength() (*big.Int, error) {
	return _MetaMorphoVault.Contract.WithdrawQueueLength(&_MetaMorphoVault.CallOpts)
}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_MetaMorphoVault *MetaMorphoVaultCallerSession) WithdrawQueueLength() (*big.Int, error) {
	return _MetaMorphoVault.Contract.WithdrawQueueLength(&_MetaMorphoVault.CallOpts)
}

// AcceptCap is a paid mutator transaction binding the contract method 0x6fda3868.
//
// Solidity: function acceptCap((address,address,address,address,uint256) marketParams) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) AcceptCap(opts *bind.TransactOpts, marketParams MarketParams) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "acceptCap", marketParams)
}

// AcceptCap is a paid mutator transaction binding the contract method 0x6fda3868.
//
// Solidity: function acceptCap((address,address,address,address,uint256) marketParams) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) AcceptCap(marketParams MarketParams) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.AcceptCap(&_MetaMorphoVault.TransactOpts, marketParams)
}

// AcceptCap is a paid mutator transaction binding the contract method 0x6fda3868.
//
// Solidity: function acceptCap((address,address,address,address,uint256) marketParams) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) AcceptCap(marketParams MarketParams) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.AcceptCap(&_MetaMorphoVault.TransactOpts, marketParams)
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) AcceptGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "acceptGuardian")
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) AcceptGuardian() (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.AcceptGuardian(&_MetaMorphoVault.TransactOpts)
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) AcceptGuardian() (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.AcceptGuardian(&_MetaMorphoVault.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) AcceptOwnership() (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.AcceptOwnership(&_MetaMorphoVault.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.AcceptOwnership(&_MetaMorphoVault.TransactOpts)
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) AcceptTimelock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "acceptTimelock")
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) AcceptTimelock() (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.AcceptTimelock(&_MetaMorphoVault.TransactOpts)
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) AcceptTimelock() (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.AcceptTimelock(&_MetaMorphoVault.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MetaMorphoVault *MetaMorphoVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MetaMorphoVault *MetaMorphoVaultSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Approve(&_MetaMorphoVault.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Approve(&_MetaMorphoVault.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_MetaMorphoVault *MetaMorphoVaultTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_MetaMorphoVault *MetaMorphoVaultSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Deposit(&_MetaMorphoVault.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Deposit(&_MetaMorphoVault.TransactOpts, assets, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_MetaMorphoVault *MetaMorphoVaultTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_MetaMorphoVault *MetaMorphoVaultSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Mint(&_MetaMorphoVault.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Mint(&_MetaMorphoVault.TransactOpts, shares, receiver)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_MetaMorphoVault *MetaMorphoVaultTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_MetaMorphoVault *MetaMorphoVaultSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Multicall(&_MetaMorphoVault.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Multicall(&_MetaMorphoVault.TransactOpts, data)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Permit(&_MetaMorphoVault.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Permit(&_MetaMorphoVault.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Reallocate is a paid mutator transaction binding the contract method 0x7299aa31.
//
// Solidity: function reallocate(((address,address,address,address,uint256),uint256)[] allocations) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) Reallocate(opts *bind.TransactOpts, allocations []MarketAllocation) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "reallocate", allocations)
}

// Reallocate is a paid mutator transaction binding the contract method 0x7299aa31.
//
// Solidity: function reallocate(((address,address,address,address,uint256),uint256)[] allocations) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) Reallocate(allocations []MarketAllocation) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Reallocate(&_MetaMorphoVault.TransactOpts, allocations)
}

// Reallocate is a paid mutator transaction binding the contract method 0x7299aa31.
//
// Solidity: function reallocate(((address,address,address,address,uint256),uint256)[] allocations) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) Reallocate(allocations []MarketAllocation) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Reallocate(&_MetaMorphoVault.TransactOpts, allocations)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_MetaMorphoVault *MetaMorphoVaultTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_MetaMorphoVault *MetaMorphoVaultSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Redeem(&_MetaMorphoVault.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Redeem(&_MetaMorphoVault.TransactOpts, shares, receiver, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.RenounceOwnership(&_MetaMorphoVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.RenounceOwnership(&_MetaMorphoVault.TransactOpts)
}

// RevokePendingCap is a paid mutator transaction binding the contract method 0x102f7b6c.
//
// Solidity: function revokePendingCap(bytes32 id) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) RevokePendingCap(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "revokePendingCap", id)
}

// RevokePendingCap is a paid mutator transaction binding the contract method 0x102f7b6c.
//
// Solidity: function revokePendingCap(bytes32 id) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) RevokePendingCap(id [32]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.RevokePendingCap(&_MetaMorphoVault.TransactOpts, id)
}

// RevokePendingCap is a paid mutator transaction binding the contract method 0x102f7b6c.
//
// Solidity: function revokePendingCap(bytes32 id) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) RevokePendingCap(id [32]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.RevokePendingCap(&_MetaMorphoVault.TransactOpts, id)
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) RevokePendingGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "revokePendingGuardian")
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) RevokePendingGuardian() (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.RevokePendingGuardian(&_MetaMorphoVault.TransactOpts)
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) RevokePendingGuardian() (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.RevokePendingGuardian(&_MetaMorphoVault.TransactOpts)
}

// RevokePendingMarketRemoval is a paid mutator transaction binding the contract method 0x4b998de5.
//
// Solidity: function revokePendingMarketRemoval(bytes32 id) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) RevokePendingMarketRemoval(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "revokePendingMarketRemoval", id)
}

// RevokePendingMarketRemoval is a paid mutator transaction binding the contract method 0x4b998de5.
//
// Solidity: function revokePendingMarketRemoval(bytes32 id) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) RevokePendingMarketRemoval(id [32]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.RevokePendingMarketRemoval(&_MetaMorphoVault.TransactOpts, id)
}

// RevokePendingMarketRemoval is a paid mutator transaction binding the contract method 0x4b998de5.
//
// Solidity: function revokePendingMarketRemoval(bytes32 id) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) RevokePendingMarketRemoval(id [32]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.RevokePendingMarketRemoval(&_MetaMorphoVault.TransactOpts, id)
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) RevokePendingTimelock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "revokePendingTimelock")
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) RevokePendingTimelock() (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.RevokePendingTimelock(&_MetaMorphoVault.TransactOpts)
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) RevokePendingTimelock() (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.RevokePendingTimelock(&_MetaMorphoVault.TransactOpts)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) SetCurator(opts *bind.TransactOpts, newCurator common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "setCurator", newCurator)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) SetCurator(newCurator common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SetCurator(&_MetaMorphoVault.TransactOpts, newCurator)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) SetCurator(newCurator common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SetCurator(&_MetaMorphoVault.TransactOpts, newCurator)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) SetFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "setFee", newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SetFee(&_MetaMorphoVault.TransactOpts, newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SetFee(&_MetaMorphoVault.TransactOpts, newFee)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) SetFeeRecipient(opts *bind.TransactOpts, newFeeRecipient common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "setFeeRecipient", newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SetFeeRecipient(&_MetaMorphoVault.TransactOpts, newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SetFeeRecipient(&_MetaMorphoVault.TransactOpts, newFeeRecipient)
}

// SetIsAllocator is a paid mutator transaction binding the contract method 0xb192a84a.
//
// Solidity: function setIsAllocator(address newAllocator, bool newIsAllocator) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) SetIsAllocator(opts *bind.TransactOpts, newAllocator common.Address, newIsAllocator bool) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "setIsAllocator", newAllocator, newIsAllocator)
}

// SetIsAllocator is a paid mutator transaction binding the contract method 0xb192a84a.
//
// Solidity: function setIsAllocator(address newAllocator, bool newIsAllocator) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) SetIsAllocator(newAllocator common.Address, newIsAllocator bool) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SetIsAllocator(&_MetaMorphoVault.TransactOpts, newAllocator, newIsAllocator)
}

// SetIsAllocator is a paid mutator transaction binding the contract method 0xb192a84a.
//
// Solidity: function setIsAllocator(address newAllocator, bool newIsAllocator) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) SetIsAllocator(newAllocator common.Address, newIsAllocator bool) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SetIsAllocator(&_MetaMorphoVault.TransactOpts, newAllocator, newIsAllocator)
}

// SetSkimRecipient is a paid mutator transaction binding the contract method 0x2b30997b.
//
// Solidity: function setSkimRecipient(address newSkimRecipient) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) SetSkimRecipient(opts *bind.TransactOpts, newSkimRecipient common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "setSkimRecipient", newSkimRecipient)
}

// SetSkimRecipient is a paid mutator transaction binding the contract method 0x2b30997b.
//
// Solidity: function setSkimRecipient(address newSkimRecipient) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) SetSkimRecipient(newSkimRecipient common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SetSkimRecipient(&_MetaMorphoVault.TransactOpts, newSkimRecipient)
}

// SetSkimRecipient is a paid mutator transaction binding the contract method 0x2b30997b.
//
// Solidity: function setSkimRecipient(address newSkimRecipient) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) SetSkimRecipient(newSkimRecipient common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SetSkimRecipient(&_MetaMorphoVault.TransactOpts, newSkimRecipient)
}

// SetSupplyQueue is a paid mutator transaction binding the contract method 0x2acc56f9.
//
// Solidity: function setSupplyQueue(bytes32[] newSupplyQueue) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) SetSupplyQueue(opts *bind.TransactOpts, newSupplyQueue [][32]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "setSupplyQueue", newSupplyQueue)
}

// SetSupplyQueue is a paid mutator transaction binding the contract method 0x2acc56f9.
//
// Solidity: function setSupplyQueue(bytes32[] newSupplyQueue) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) SetSupplyQueue(newSupplyQueue [][32]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SetSupplyQueue(&_MetaMorphoVault.TransactOpts, newSupplyQueue)
}

// SetSupplyQueue is a paid mutator transaction binding the contract method 0x2acc56f9.
//
// Solidity: function setSupplyQueue(bytes32[] newSupplyQueue) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) SetSupplyQueue(newSupplyQueue [][32]byte) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SetSupplyQueue(&_MetaMorphoVault.TransactOpts, newSupplyQueue)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) Skim(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "skim", token)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) Skim(token common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Skim(&_MetaMorphoVault.TransactOpts, token)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) Skim(token common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Skim(&_MetaMorphoVault.TransactOpts, token)
}

// SubmitCap is a paid mutator transaction binding the contract method 0x3b24c2bf.
//
// Solidity: function submitCap((address,address,address,address,uint256) marketParams, uint256 newSupplyCap) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) SubmitCap(opts *bind.TransactOpts, marketParams MarketParams, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "submitCap", marketParams, newSupplyCap)
}

// SubmitCap is a paid mutator transaction binding the contract method 0x3b24c2bf.
//
// Solidity: function submitCap((address,address,address,address,uint256) marketParams, uint256 newSupplyCap) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) SubmitCap(marketParams MarketParams, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SubmitCap(&_MetaMorphoVault.TransactOpts, marketParams, newSupplyCap)
}

// SubmitCap is a paid mutator transaction binding the contract method 0x3b24c2bf.
//
// Solidity: function submitCap((address,address,address,address,uint256) marketParams, uint256 newSupplyCap) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) SubmitCap(marketParams MarketParams, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SubmitCap(&_MetaMorphoVault.TransactOpts, marketParams, newSupplyCap)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) SubmitGuardian(opts *bind.TransactOpts, newGuardian common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "submitGuardian", newGuardian)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) SubmitGuardian(newGuardian common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SubmitGuardian(&_MetaMorphoVault.TransactOpts, newGuardian)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) SubmitGuardian(newGuardian common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SubmitGuardian(&_MetaMorphoVault.TransactOpts, newGuardian)
}

// SubmitMarketRemoval is a paid mutator transaction binding the contract method 0x84755b5f.
//
// Solidity: function submitMarketRemoval((address,address,address,address,uint256) marketParams) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) SubmitMarketRemoval(opts *bind.TransactOpts, marketParams MarketParams) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "submitMarketRemoval", marketParams)
}

// SubmitMarketRemoval is a paid mutator transaction binding the contract method 0x84755b5f.
//
// Solidity: function submitMarketRemoval((address,address,address,address,uint256) marketParams) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) SubmitMarketRemoval(marketParams MarketParams) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SubmitMarketRemoval(&_MetaMorphoVault.TransactOpts, marketParams)
}

// SubmitMarketRemoval is a paid mutator transaction binding the contract method 0x84755b5f.
//
// Solidity: function submitMarketRemoval((address,address,address,address,uint256) marketParams) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) SubmitMarketRemoval(marketParams MarketParams) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SubmitMarketRemoval(&_MetaMorphoVault.TransactOpts, marketParams)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) SubmitTimelock(opts *bind.TransactOpts, newTimelock *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "submitTimelock", newTimelock)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) SubmitTimelock(newTimelock *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SubmitTimelock(&_MetaMorphoVault.TransactOpts, newTimelock)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) SubmitTimelock(newTimelock *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.SubmitTimelock(&_MetaMorphoVault.TransactOpts, newTimelock)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MetaMorphoVault *MetaMorphoVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MetaMorphoVault *MetaMorphoVaultSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Transfer(&_MetaMorphoVault.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Transfer(&_MetaMorphoVault.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MetaMorphoVault *MetaMorphoVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MetaMorphoVault *MetaMorphoVaultSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.TransferFrom(&_MetaMorphoVault.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.TransferFrom(&_MetaMorphoVault.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.TransferOwnership(&_MetaMorphoVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.TransferOwnership(&_MetaMorphoVault.TransactOpts, newOwner)
}

// UpdateWithdrawQueue is a paid mutator transaction binding the contract method 0x41b67833.
//
// Solidity: function updateWithdrawQueue(uint256[] indexes) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactor) UpdateWithdrawQueue(opts *bind.TransactOpts, indexes []*big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "updateWithdrawQueue", indexes)
}

// UpdateWithdrawQueue is a paid mutator transaction binding the contract method 0x41b67833.
//
// Solidity: function updateWithdrawQueue(uint256[] indexes) returns()
func (_MetaMorphoVault *MetaMorphoVaultSession) UpdateWithdrawQueue(indexes []*big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.UpdateWithdrawQueue(&_MetaMorphoVault.TransactOpts, indexes)
}

// UpdateWithdrawQueue is a paid mutator transaction binding the contract method 0x41b67833.
//
// Solidity: function updateWithdrawQueue(uint256[] indexes) returns()
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) UpdateWithdrawQueue(indexes []*big.Int) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.UpdateWithdrawQueue(&_MetaMorphoVault.TransactOpts, indexes)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_MetaMorphoVault *MetaMorphoVaultTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_MetaMorphoVault *MetaMorphoVaultSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Withdraw(&_MetaMorphoVault.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_MetaMorphoVault *MetaMorphoVaultTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _MetaMorphoVault.Contract.Withdraw(&_MetaMorphoVault.TransactOpts, assets, receiver, owner)
}

// MetaMorphoVaultAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the MetaMorphoVault contract.
type MetaMorphoVaultAccrueInterestIterator struct {
	Event *MetaMorphoVaultAccrueInterest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultAccrueInterest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultAccrueInterest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultAccrueInterest represents a AccrueInterest event raised by the MetaMorphoVault contract.
type MetaMorphoVaultAccrueInterest struct {
	NewTotalAssets *big.Int
	FeeShares      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0xf66f28b40975dbb933913542c7e6a0f50a1d0f20aa74ea6e0efe65ab616323ec.
//
// Solidity: event AccrueInterest(uint256 newTotalAssets, uint256 feeShares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*MetaMorphoVaultAccrueInterestIterator, error) {

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultAccrueInterestIterator{contract: _MetaMorphoVault.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0xf66f28b40975dbb933913542c7e6a0f50a1d0f20aa74ea6e0efe65ab616323ec.
//
// Solidity: event AccrueInterest(uint256 newTotalAssets, uint256 feeShares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultAccrueInterest)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccrueInterest is a log parse operation binding the contract event 0xf66f28b40975dbb933913542c7e6a0f50a1d0f20aa74ea6e0efe65ab616323ec.
//
// Solidity: event AccrueInterest(uint256 newTotalAssets, uint256 feeShares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseAccrueInterest(log types.Log) (*MetaMorphoVaultAccrueInterest, error) {
	event := new(MetaMorphoVaultAccrueInterest)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MetaMorphoVault contract.
type MetaMorphoVaultApprovalIterator struct {
	Event *MetaMorphoVaultApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultApproval represents a Approval event raised by the MetaMorphoVault contract.
type MetaMorphoVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MetaMorphoVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultApprovalIterator{contract: _MetaMorphoVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultApproval)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseApproval(log types.Log) (*MetaMorphoVaultApproval, error) {
	event := new(MetaMorphoVaultApproval)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MetaMorphoVault contract.
type MetaMorphoVaultDepositIterator struct {
	Event *MetaMorphoVaultDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultDeposit represents a Deposit event raised by the MetaMorphoVault contract.
type MetaMorphoVaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*MetaMorphoVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultDepositIterator{contract: _MetaMorphoVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultDeposit)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseDeposit(log types.Log) (*MetaMorphoVaultDeposit, error) {
	event := new(MetaMorphoVaultDeposit)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the MetaMorphoVault contract.
type MetaMorphoVaultEIP712DomainChangedIterator struct {
	Event *MetaMorphoVaultEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultEIP712DomainChanged represents a EIP712DomainChanged event raised by the MetaMorphoVault contract.
type MetaMorphoVaultEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*MetaMorphoVaultEIP712DomainChangedIterator, error) {

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultEIP712DomainChangedIterator{contract: _MetaMorphoVault.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultEIP712DomainChanged)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseEIP712DomainChanged(log types.Log) (*MetaMorphoVaultEIP712DomainChanged, error) {
	event := new(MetaMorphoVaultEIP712DomainChanged)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the MetaMorphoVault contract.
type MetaMorphoVaultOwnershipTransferStartedIterator struct {
	Event *MetaMorphoVaultOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the MetaMorphoVault contract.
type MetaMorphoVaultOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MetaMorphoVaultOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultOwnershipTransferStartedIterator{contract: _MetaMorphoVault.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultOwnershipTransferStarted)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseOwnershipTransferStarted(log types.Log) (*MetaMorphoVaultOwnershipTransferStarted, error) {
	event := new(MetaMorphoVaultOwnershipTransferStarted)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MetaMorphoVault contract.
type MetaMorphoVaultOwnershipTransferredIterator struct {
	Event *MetaMorphoVaultOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultOwnershipTransferred represents a OwnershipTransferred event raised by the MetaMorphoVault contract.
type MetaMorphoVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MetaMorphoVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultOwnershipTransferredIterator{contract: _MetaMorphoVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultOwnershipTransferred)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseOwnershipTransferred(log types.Log) (*MetaMorphoVaultOwnershipTransferred, error) {
	event := new(MetaMorphoVaultOwnershipTransferred)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultReallocateSupplyIterator is returned from FilterReallocateSupply and is used to iterate over the raw logs and unpacked data for ReallocateSupply events raised by the MetaMorphoVault contract.
type MetaMorphoVaultReallocateSupplyIterator struct {
	Event *MetaMorphoVaultReallocateSupply // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultReallocateSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultReallocateSupply)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultReallocateSupply)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultReallocateSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultReallocateSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultReallocateSupply represents a ReallocateSupply event raised by the MetaMorphoVault contract.
type MetaMorphoVaultReallocateSupply struct {
	Caller         common.Address
	Id             [32]byte
	SuppliedAssets *big.Int
	SuppliedShares *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReallocateSupply is a free log retrieval operation binding the contract event 0x89bf199df65bf65155e3e0a8abc4ad4a1be606220c8295840dba2ab5656c1f6d.
//
// Solidity: event ReallocateSupply(address indexed caller, bytes32 indexed id, uint256 suppliedAssets, uint256 suppliedShares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterReallocateSupply(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MetaMorphoVaultReallocateSupplyIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "ReallocateSupply", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultReallocateSupplyIterator{contract: _MetaMorphoVault.contract, event: "ReallocateSupply", logs: logs, sub: sub}, nil
}

// WatchReallocateSupply is a free log subscription operation binding the contract event 0x89bf199df65bf65155e3e0a8abc4ad4a1be606220c8295840dba2ab5656c1f6d.
//
// Solidity: event ReallocateSupply(address indexed caller, bytes32 indexed id, uint256 suppliedAssets, uint256 suppliedShares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchReallocateSupply(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultReallocateSupply, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "ReallocateSupply", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultReallocateSupply)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "ReallocateSupply", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReallocateSupply is a log parse operation binding the contract event 0x89bf199df65bf65155e3e0a8abc4ad4a1be606220c8295840dba2ab5656c1f6d.
//
// Solidity: event ReallocateSupply(address indexed caller, bytes32 indexed id, uint256 suppliedAssets, uint256 suppliedShares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseReallocateSupply(log types.Log) (*MetaMorphoVaultReallocateSupply, error) {
	event := new(MetaMorphoVaultReallocateSupply)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "ReallocateSupply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultReallocateWithdrawIterator is returned from FilterReallocateWithdraw and is used to iterate over the raw logs and unpacked data for ReallocateWithdraw events raised by the MetaMorphoVault contract.
type MetaMorphoVaultReallocateWithdrawIterator struct {
	Event *MetaMorphoVaultReallocateWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultReallocateWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultReallocateWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultReallocateWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultReallocateWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultReallocateWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultReallocateWithdraw represents a ReallocateWithdraw event raised by the MetaMorphoVault contract.
type MetaMorphoVaultReallocateWithdraw struct {
	Caller          common.Address
	Id              [32]byte
	WithdrawnAssets *big.Int
	WithdrawnShares *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReallocateWithdraw is a free log retrieval operation binding the contract event 0xdd8bf5226dff861316e0fa7863fdb7dc7b87c614eb29a135f524eb79d5a1189a.
//
// Solidity: event ReallocateWithdraw(address indexed caller, bytes32 indexed id, uint256 withdrawnAssets, uint256 withdrawnShares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterReallocateWithdraw(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MetaMorphoVaultReallocateWithdrawIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "ReallocateWithdraw", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultReallocateWithdrawIterator{contract: _MetaMorphoVault.contract, event: "ReallocateWithdraw", logs: logs, sub: sub}, nil
}

// WatchReallocateWithdraw is a free log subscription operation binding the contract event 0xdd8bf5226dff861316e0fa7863fdb7dc7b87c614eb29a135f524eb79d5a1189a.
//
// Solidity: event ReallocateWithdraw(address indexed caller, bytes32 indexed id, uint256 withdrawnAssets, uint256 withdrawnShares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchReallocateWithdraw(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultReallocateWithdraw, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "ReallocateWithdraw", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultReallocateWithdraw)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "ReallocateWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReallocateWithdraw is a log parse operation binding the contract event 0xdd8bf5226dff861316e0fa7863fdb7dc7b87c614eb29a135f524eb79d5a1189a.
//
// Solidity: event ReallocateWithdraw(address indexed caller, bytes32 indexed id, uint256 withdrawnAssets, uint256 withdrawnShares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseReallocateWithdraw(log types.Log) (*MetaMorphoVaultReallocateWithdraw, error) {
	event := new(MetaMorphoVaultReallocateWithdraw)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "ReallocateWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultRevokePendingCapIterator is returned from FilterRevokePendingCap and is used to iterate over the raw logs and unpacked data for RevokePendingCap events raised by the MetaMorphoVault contract.
type MetaMorphoVaultRevokePendingCapIterator struct {
	Event *MetaMorphoVaultRevokePendingCap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultRevokePendingCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultRevokePendingCap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultRevokePendingCap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultRevokePendingCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultRevokePendingCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultRevokePendingCap represents a RevokePendingCap event raised by the MetaMorphoVault contract.
type MetaMorphoVaultRevokePendingCap struct {
	Caller common.Address
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingCap is a free log retrieval operation binding the contract event 0x1026ceca5ed3747eb5edec555732d4a6f901ce1a875ecf981064628cadde1120.
//
// Solidity: event RevokePendingCap(address indexed caller, bytes32 indexed id)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterRevokePendingCap(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MetaMorphoVaultRevokePendingCapIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "RevokePendingCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultRevokePendingCapIterator{contract: _MetaMorphoVault.contract, event: "RevokePendingCap", logs: logs, sub: sub}, nil
}

// WatchRevokePendingCap is a free log subscription operation binding the contract event 0x1026ceca5ed3747eb5edec555732d4a6f901ce1a875ecf981064628cadde1120.
//
// Solidity: event RevokePendingCap(address indexed caller, bytes32 indexed id)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchRevokePendingCap(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultRevokePendingCap, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "RevokePendingCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultRevokePendingCap)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "RevokePendingCap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokePendingCap is a log parse operation binding the contract event 0x1026ceca5ed3747eb5edec555732d4a6f901ce1a875ecf981064628cadde1120.
//
// Solidity: event RevokePendingCap(address indexed caller, bytes32 indexed id)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseRevokePendingCap(log types.Log) (*MetaMorphoVaultRevokePendingCap, error) {
	event := new(MetaMorphoVaultRevokePendingCap)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "RevokePendingCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultRevokePendingGuardianIterator is returned from FilterRevokePendingGuardian and is used to iterate over the raw logs and unpacked data for RevokePendingGuardian events raised by the MetaMorphoVault contract.
type MetaMorphoVaultRevokePendingGuardianIterator struct {
	Event *MetaMorphoVaultRevokePendingGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultRevokePendingGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultRevokePendingGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultRevokePendingGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultRevokePendingGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultRevokePendingGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultRevokePendingGuardian represents a RevokePendingGuardian event raised by the MetaMorphoVault contract.
type MetaMorphoVaultRevokePendingGuardian struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingGuardian is a free log retrieval operation binding the contract event 0xc40a085ccfa20f5fd518ade5c3a77a7ecbdfbb4c75efcdca6146a8e3c841d663.
//
// Solidity: event RevokePendingGuardian(address indexed caller)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterRevokePendingGuardian(opts *bind.FilterOpts, caller []common.Address) (*MetaMorphoVaultRevokePendingGuardianIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "RevokePendingGuardian", callerRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultRevokePendingGuardianIterator{contract: _MetaMorphoVault.contract, event: "RevokePendingGuardian", logs: logs, sub: sub}, nil
}

// WatchRevokePendingGuardian is a free log subscription operation binding the contract event 0xc40a085ccfa20f5fd518ade5c3a77a7ecbdfbb4c75efcdca6146a8e3c841d663.
//
// Solidity: event RevokePendingGuardian(address indexed caller)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchRevokePendingGuardian(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultRevokePendingGuardian, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "RevokePendingGuardian", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultRevokePendingGuardian)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "RevokePendingGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokePendingGuardian is a log parse operation binding the contract event 0xc40a085ccfa20f5fd518ade5c3a77a7ecbdfbb4c75efcdca6146a8e3c841d663.
//
// Solidity: event RevokePendingGuardian(address indexed caller)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseRevokePendingGuardian(log types.Log) (*MetaMorphoVaultRevokePendingGuardian, error) {
	event := new(MetaMorphoVaultRevokePendingGuardian)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "RevokePendingGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultRevokePendingMarketRemovalIterator is returned from FilterRevokePendingMarketRemoval and is used to iterate over the raw logs and unpacked data for RevokePendingMarketRemoval events raised by the MetaMorphoVault contract.
type MetaMorphoVaultRevokePendingMarketRemovalIterator struct {
	Event *MetaMorphoVaultRevokePendingMarketRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultRevokePendingMarketRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultRevokePendingMarketRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultRevokePendingMarketRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultRevokePendingMarketRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultRevokePendingMarketRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultRevokePendingMarketRemoval represents a RevokePendingMarketRemoval event raised by the MetaMorphoVault contract.
type MetaMorphoVaultRevokePendingMarketRemoval struct {
	Caller common.Address
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingMarketRemoval is a free log retrieval operation binding the contract event 0xcbeb8ecdaa5a3c133e62219b63bfc35bce3fda13065d2bed32e3b7dde60a59f4.
//
// Solidity: event RevokePendingMarketRemoval(address indexed caller, bytes32 indexed id)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterRevokePendingMarketRemoval(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MetaMorphoVaultRevokePendingMarketRemovalIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "RevokePendingMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultRevokePendingMarketRemovalIterator{contract: _MetaMorphoVault.contract, event: "RevokePendingMarketRemoval", logs: logs, sub: sub}, nil
}

// WatchRevokePendingMarketRemoval is a free log subscription operation binding the contract event 0xcbeb8ecdaa5a3c133e62219b63bfc35bce3fda13065d2bed32e3b7dde60a59f4.
//
// Solidity: event RevokePendingMarketRemoval(address indexed caller, bytes32 indexed id)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchRevokePendingMarketRemoval(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultRevokePendingMarketRemoval, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "RevokePendingMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultRevokePendingMarketRemoval)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "RevokePendingMarketRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokePendingMarketRemoval is a log parse operation binding the contract event 0xcbeb8ecdaa5a3c133e62219b63bfc35bce3fda13065d2bed32e3b7dde60a59f4.
//
// Solidity: event RevokePendingMarketRemoval(address indexed caller, bytes32 indexed id)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseRevokePendingMarketRemoval(log types.Log) (*MetaMorphoVaultRevokePendingMarketRemoval, error) {
	event := new(MetaMorphoVaultRevokePendingMarketRemoval)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "RevokePendingMarketRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultRevokePendingTimelockIterator is returned from FilterRevokePendingTimelock and is used to iterate over the raw logs and unpacked data for RevokePendingTimelock events raised by the MetaMorphoVault contract.
type MetaMorphoVaultRevokePendingTimelockIterator struct {
	Event *MetaMorphoVaultRevokePendingTimelock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultRevokePendingTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultRevokePendingTimelock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultRevokePendingTimelock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultRevokePendingTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultRevokePendingTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultRevokePendingTimelock represents a RevokePendingTimelock event raised by the MetaMorphoVault contract.
type MetaMorphoVaultRevokePendingTimelock struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingTimelock is a free log retrieval operation binding the contract event 0x921828337692c347c634c5d2aacbc7b756014674bd236f3cc2058d8e284a951b.
//
// Solidity: event RevokePendingTimelock(address indexed caller)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterRevokePendingTimelock(opts *bind.FilterOpts, caller []common.Address) (*MetaMorphoVaultRevokePendingTimelockIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "RevokePendingTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultRevokePendingTimelockIterator{contract: _MetaMorphoVault.contract, event: "RevokePendingTimelock", logs: logs, sub: sub}, nil
}

// WatchRevokePendingTimelock is a free log subscription operation binding the contract event 0x921828337692c347c634c5d2aacbc7b756014674bd236f3cc2058d8e284a951b.
//
// Solidity: event RevokePendingTimelock(address indexed caller)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchRevokePendingTimelock(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultRevokePendingTimelock, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "RevokePendingTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultRevokePendingTimelock)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "RevokePendingTimelock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokePendingTimelock is a log parse operation binding the contract event 0x921828337692c347c634c5d2aacbc7b756014674bd236f3cc2058d8e284a951b.
//
// Solidity: event RevokePendingTimelock(address indexed caller)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseRevokePendingTimelock(log types.Log) (*MetaMorphoVaultRevokePendingTimelock, error) {
	event := new(MetaMorphoVaultRevokePendingTimelock)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "RevokePendingTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSetCapIterator is returned from FilterSetCap and is used to iterate over the raw logs and unpacked data for SetCap events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetCapIterator struct {
	Event *MetaMorphoVaultSetCap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSetCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSetCap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSetCap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSetCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSetCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSetCap represents a SetCap event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetCap struct {
	Caller common.Address
	Id     [32]byte
	Cap    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetCap is a free log retrieval operation binding the contract event 0xe86b6d3313d3098f4c5f689c935de8fde876a597c185def2cedab85efedac686.
//
// Solidity: event SetCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSetCap(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MetaMorphoVaultSetCapIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "SetCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSetCapIterator{contract: _MetaMorphoVault.contract, event: "SetCap", logs: logs, sub: sub}, nil
}

// WatchSetCap is a free log subscription operation binding the contract event 0xe86b6d3313d3098f4c5f689c935de8fde876a597c185def2cedab85efedac686.
//
// Solidity: event SetCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSetCap(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSetCap, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "SetCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSetCap)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "SetCap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetCap is a log parse operation binding the contract event 0xe86b6d3313d3098f4c5f689c935de8fde876a597c185def2cedab85efedac686.
//
// Solidity: event SetCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSetCap(log types.Log) (*MetaMorphoVaultSetCap, error) {
	event := new(MetaMorphoVaultSetCap)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "SetCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSetCuratorIterator is returned from FilterSetCurator and is used to iterate over the raw logs and unpacked data for SetCurator events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetCuratorIterator struct {
	Event *MetaMorphoVaultSetCurator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSetCuratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSetCurator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSetCurator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSetCuratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSetCuratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSetCurator represents a SetCurator event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetCurator struct {
	NewCurator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetCurator is a free log retrieval operation binding the contract event 0xbd0a63c12948fbc9194a5839019f99c9d71db924e5c70018265bc778b8f1a506.
//
// Solidity: event SetCurator(address indexed newCurator)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSetCurator(opts *bind.FilterOpts, newCurator []common.Address) (*MetaMorphoVaultSetCuratorIterator, error) {

	var newCuratorRule []interface{}
	for _, newCuratorItem := range newCurator {
		newCuratorRule = append(newCuratorRule, newCuratorItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "SetCurator", newCuratorRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSetCuratorIterator{contract: _MetaMorphoVault.contract, event: "SetCurator", logs: logs, sub: sub}, nil
}

// WatchSetCurator is a free log subscription operation binding the contract event 0xbd0a63c12948fbc9194a5839019f99c9d71db924e5c70018265bc778b8f1a506.
//
// Solidity: event SetCurator(address indexed newCurator)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSetCurator(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSetCurator, newCurator []common.Address) (event.Subscription, error) {

	var newCuratorRule []interface{}
	for _, newCuratorItem := range newCurator {
		newCuratorRule = append(newCuratorRule, newCuratorItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "SetCurator", newCuratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSetCurator)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "SetCurator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetCurator is a log parse operation binding the contract event 0xbd0a63c12948fbc9194a5839019f99c9d71db924e5c70018265bc778b8f1a506.
//
// Solidity: event SetCurator(address indexed newCurator)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSetCurator(log types.Log) (*MetaMorphoVaultSetCurator, error) {
	event := new(MetaMorphoVaultSetCurator)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "SetCurator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetFeeIterator struct {
	Event *MetaMorphoVaultSetFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSetFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSetFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSetFee represents a SetFee event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetFee struct {
	Caller common.Address
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x01fe2943baee27f47add82886c2200f910c749c461c9b63c5fe83901a53bdb49.
//
// Solidity: event SetFee(address indexed caller, uint256 newFee)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSetFee(opts *bind.FilterOpts, caller []common.Address) (*MetaMorphoVaultSetFeeIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "SetFee", callerRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSetFeeIterator{contract: _MetaMorphoVault.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x01fe2943baee27f47add82886c2200f910c749c461c9b63c5fe83901a53bdb49.
//
// Solidity: event SetFee(address indexed caller, uint256 newFee)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSetFee, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "SetFee", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSetFee)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "SetFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFee is a log parse operation binding the contract event 0x01fe2943baee27f47add82886c2200f910c749c461c9b63c5fe83901a53bdb49.
//
// Solidity: event SetFee(address indexed caller, uint256 newFee)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSetFee(log types.Log) (*MetaMorphoVaultSetFee, error) {
	event := new(MetaMorphoVaultSetFee)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSetFeeRecipientIterator is returned from FilterSetFeeRecipient and is used to iterate over the raw logs and unpacked data for SetFeeRecipient events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetFeeRecipientIterator struct {
	Event *MetaMorphoVaultSetFeeRecipient // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSetFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSetFeeRecipient)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSetFeeRecipient)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSetFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSetFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSetFeeRecipient represents a SetFeeRecipient event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetFeeRecipient struct {
	NewFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetFeeRecipient is a free log retrieval operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address indexed newFeeRecipient)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSetFeeRecipient(opts *bind.FilterOpts, newFeeRecipient []common.Address) (*MetaMorphoVaultSetFeeRecipientIterator, error) {

	var newFeeRecipientRule []interface{}
	for _, newFeeRecipientItem := range newFeeRecipient {
		newFeeRecipientRule = append(newFeeRecipientRule, newFeeRecipientItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "SetFeeRecipient", newFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSetFeeRecipientIterator{contract: _MetaMorphoVault.contract, event: "SetFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchSetFeeRecipient is a free log subscription operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address indexed newFeeRecipient)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSetFeeRecipient(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSetFeeRecipient, newFeeRecipient []common.Address) (event.Subscription, error) {

	var newFeeRecipientRule []interface{}
	for _, newFeeRecipientItem := range newFeeRecipient {
		newFeeRecipientRule = append(newFeeRecipientRule, newFeeRecipientItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "SetFeeRecipient", newFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSetFeeRecipient)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "SetFeeRecipient", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFeeRecipient is a log parse operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address indexed newFeeRecipient)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSetFeeRecipient(log types.Log) (*MetaMorphoVaultSetFeeRecipient, error) {
	event := new(MetaMorphoVaultSetFeeRecipient)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "SetFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSetGuardianIterator is returned from FilterSetGuardian and is used to iterate over the raw logs and unpacked data for SetGuardian events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetGuardianIterator struct {
	Event *MetaMorphoVaultSetGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSetGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSetGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSetGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSetGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSetGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSetGuardian represents a SetGuardian event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetGuardian struct {
	Caller   common.Address
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetGuardian is a free log retrieval operation binding the contract event 0xcb11cc8aade2f5a556749d1b2380d108a16fac3431e6a5d5ce12ef9de0bd76e3.
//
// Solidity: event SetGuardian(address indexed caller, address indexed guardian)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSetGuardian(opts *bind.FilterOpts, caller []common.Address, guardian []common.Address) (*MetaMorphoVaultSetGuardianIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "SetGuardian", callerRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSetGuardianIterator{contract: _MetaMorphoVault.contract, event: "SetGuardian", logs: logs, sub: sub}, nil
}

// WatchSetGuardian is a free log subscription operation binding the contract event 0xcb11cc8aade2f5a556749d1b2380d108a16fac3431e6a5d5ce12ef9de0bd76e3.
//
// Solidity: event SetGuardian(address indexed caller, address indexed guardian)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSetGuardian(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSetGuardian, caller []common.Address, guardian []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "SetGuardian", callerRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSetGuardian)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "SetGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetGuardian is a log parse operation binding the contract event 0xcb11cc8aade2f5a556749d1b2380d108a16fac3431e6a5d5ce12ef9de0bd76e3.
//
// Solidity: event SetGuardian(address indexed caller, address indexed guardian)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSetGuardian(log types.Log) (*MetaMorphoVaultSetGuardian, error) {
	event := new(MetaMorphoVaultSetGuardian)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "SetGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSetIsAllocatorIterator is returned from FilterSetIsAllocator and is used to iterate over the raw logs and unpacked data for SetIsAllocator events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetIsAllocatorIterator struct {
	Event *MetaMorphoVaultSetIsAllocator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSetIsAllocatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSetIsAllocator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSetIsAllocator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSetIsAllocatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSetIsAllocatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSetIsAllocator represents a SetIsAllocator event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetIsAllocator struct {
	Allocator   common.Address
	IsAllocator bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetIsAllocator is a free log retrieval operation binding the contract event 0x74dc60cbc81a9472d04ad1d20e151d369c41104d655ed3f2f3091166a502cd8d.
//
// Solidity: event SetIsAllocator(address indexed allocator, bool isAllocator)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSetIsAllocator(opts *bind.FilterOpts, allocator []common.Address) (*MetaMorphoVaultSetIsAllocatorIterator, error) {

	var allocatorRule []interface{}
	for _, allocatorItem := range allocator {
		allocatorRule = append(allocatorRule, allocatorItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "SetIsAllocator", allocatorRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSetIsAllocatorIterator{contract: _MetaMorphoVault.contract, event: "SetIsAllocator", logs: logs, sub: sub}, nil
}

// WatchSetIsAllocator is a free log subscription operation binding the contract event 0x74dc60cbc81a9472d04ad1d20e151d369c41104d655ed3f2f3091166a502cd8d.
//
// Solidity: event SetIsAllocator(address indexed allocator, bool isAllocator)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSetIsAllocator(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSetIsAllocator, allocator []common.Address) (event.Subscription, error) {

	var allocatorRule []interface{}
	for _, allocatorItem := range allocator {
		allocatorRule = append(allocatorRule, allocatorItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "SetIsAllocator", allocatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSetIsAllocator)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "SetIsAllocator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetIsAllocator is a log parse operation binding the contract event 0x74dc60cbc81a9472d04ad1d20e151d369c41104d655ed3f2f3091166a502cd8d.
//
// Solidity: event SetIsAllocator(address indexed allocator, bool isAllocator)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSetIsAllocator(log types.Log) (*MetaMorphoVaultSetIsAllocator, error) {
	event := new(MetaMorphoVaultSetIsAllocator)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "SetIsAllocator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSetSkimRecipientIterator is returned from FilterSetSkimRecipient and is used to iterate over the raw logs and unpacked data for SetSkimRecipient events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetSkimRecipientIterator struct {
	Event *MetaMorphoVaultSetSkimRecipient // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSetSkimRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSetSkimRecipient)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSetSkimRecipient)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSetSkimRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSetSkimRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSetSkimRecipient represents a SetSkimRecipient event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetSkimRecipient struct {
	NewSkimRecipient common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetSkimRecipient is a free log retrieval operation binding the contract event 0x2e7908865670e21b9779422cadf5f1cba271a62bb95c71eaaf615c0a1c48ebee.
//
// Solidity: event SetSkimRecipient(address indexed newSkimRecipient)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSetSkimRecipient(opts *bind.FilterOpts, newSkimRecipient []common.Address) (*MetaMorphoVaultSetSkimRecipientIterator, error) {

	var newSkimRecipientRule []interface{}
	for _, newSkimRecipientItem := range newSkimRecipient {
		newSkimRecipientRule = append(newSkimRecipientRule, newSkimRecipientItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "SetSkimRecipient", newSkimRecipientRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSetSkimRecipientIterator{contract: _MetaMorphoVault.contract, event: "SetSkimRecipient", logs: logs, sub: sub}, nil
}

// WatchSetSkimRecipient is a free log subscription operation binding the contract event 0x2e7908865670e21b9779422cadf5f1cba271a62bb95c71eaaf615c0a1c48ebee.
//
// Solidity: event SetSkimRecipient(address indexed newSkimRecipient)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSetSkimRecipient(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSetSkimRecipient, newSkimRecipient []common.Address) (event.Subscription, error) {

	var newSkimRecipientRule []interface{}
	for _, newSkimRecipientItem := range newSkimRecipient {
		newSkimRecipientRule = append(newSkimRecipientRule, newSkimRecipientItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "SetSkimRecipient", newSkimRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSetSkimRecipient)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "SetSkimRecipient", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetSkimRecipient is a log parse operation binding the contract event 0x2e7908865670e21b9779422cadf5f1cba271a62bb95c71eaaf615c0a1c48ebee.
//
// Solidity: event SetSkimRecipient(address indexed newSkimRecipient)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSetSkimRecipient(log types.Log) (*MetaMorphoVaultSetSkimRecipient, error) {
	event := new(MetaMorphoVaultSetSkimRecipient)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "SetSkimRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSetSupplyQueueIterator is returned from FilterSetSupplyQueue and is used to iterate over the raw logs and unpacked data for SetSupplyQueue events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetSupplyQueueIterator struct {
	Event *MetaMorphoVaultSetSupplyQueue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSetSupplyQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSetSupplyQueue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSetSupplyQueue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSetSupplyQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSetSupplyQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSetSupplyQueue represents a SetSupplyQueue event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetSupplyQueue struct {
	Caller         common.Address
	NewSupplyQueue [][32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetSupplyQueue is a free log retrieval operation binding the contract event 0x6ce31538fc7fba95714ddc8a275a09252b4b1fb8f33d2550aa58a5f62ad934de.
//
// Solidity: event SetSupplyQueue(address indexed caller, bytes32[] newSupplyQueue)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSetSupplyQueue(opts *bind.FilterOpts, caller []common.Address) (*MetaMorphoVaultSetSupplyQueueIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "SetSupplyQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSetSupplyQueueIterator{contract: _MetaMorphoVault.contract, event: "SetSupplyQueue", logs: logs, sub: sub}, nil
}

// WatchSetSupplyQueue is a free log subscription operation binding the contract event 0x6ce31538fc7fba95714ddc8a275a09252b4b1fb8f33d2550aa58a5f62ad934de.
//
// Solidity: event SetSupplyQueue(address indexed caller, bytes32[] newSupplyQueue)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSetSupplyQueue(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSetSupplyQueue, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "SetSupplyQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSetSupplyQueue)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "SetSupplyQueue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetSupplyQueue is a log parse operation binding the contract event 0x6ce31538fc7fba95714ddc8a275a09252b4b1fb8f33d2550aa58a5f62ad934de.
//
// Solidity: event SetSupplyQueue(address indexed caller, bytes32[] newSupplyQueue)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSetSupplyQueue(log types.Log) (*MetaMorphoVaultSetSupplyQueue, error) {
	event := new(MetaMorphoVaultSetSupplyQueue)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "SetSupplyQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSetTimelockIterator is returned from FilterSetTimelock and is used to iterate over the raw logs and unpacked data for SetTimelock events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetTimelockIterator struct {
	Event *MetaMorphoVaultSetTimelock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSetTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSetTimelock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSetTimelock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSetTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSetTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSetTimelock represents a SetTimelock event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetTimelock struct {
	Caller      common.Address
	NewTimelock *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetTimelock is a free log retrieval operation binding the contract event 0xd28e9b90ee9b37c5936ff84392d71f29ff18117d7e76bcee60615262a90a3f75.
//
// Solidity: event SetTimelock(address indexed caller, uint256 newTimelock)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSetTimelock(opts *bind.FilterOpts, caller []common.Address) (*MetaMorphoVaultSetTimelockIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "SetTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSetTimelockIterator{contract: _MetaMorphoVault.contract, event: "SetTimelock", logs: logs, sub: sub}, nil
}

// WatchSetTimelock is a free log subscription operation binding the contract event 0xd28e9b90ee9b37c5936ff84392d71f29ff18117d7e76bcee60615262a90a3f75.
//
// Solidity: event SetTimelock(address indexed caller, uint256 newTimelock)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSetTimelock(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSetTimelock, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "SetTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSetTimelock)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "SetTimelock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTimelock is a log parse operation binding the contract event 0xd28e9b90ee9b37c5936ff84392d71f29ff18117d7e76bcee60615262a90a3f75.
//
// Solidity: event SetTimelock(address indexed caller, uint256 newTimelock)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSetTimelock(log types.Log) (*MetaMorphoVaultSetTimelock, error) {
	event := new(MetaMorphoVaultSetTimelock)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "SetTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSetWithdrawQueueIterator is returned from FilterSetWithdrawQueue and is used to iterate over the raw logs and unpacked data for SetWithdrawQueue events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetWithdrawQueueIterator struct {
	Event *MetaMorphoVaultSetWithdrawQueue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSetWithdrawQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSetWithdrawQueue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSetWithdrawQueue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSetWithdrawQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSetWithdrawQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSetWithdrawQueue represents a SetWithdrawQueue event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSetWithdrawQueue struct {
	Caller           common.Address
	NewWithdrawQueue [][32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetWithdrawQueue is a free log retrieval operation binding the contract event 0xe0c2db6b54586be6d7d49943139fccf0dd315ba63e55364a76c73cd8fdba724d.
//
// Solidity: event SetWithdrawQueue(address indexed caller, bytes32[] newWithdrawQueue)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSetWithdrawQueue(opts *bind.FilterOpts, caller []common.Address) (*MetaMorphoVaultSetWithdrawQueueIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "SetWithdrawQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSetWithdrawQueueIterator{contract: _MetaMorphoVault.contract, event: "SetWithdrawQueue", logs: logs, sub: sub}, nil
}

// WatchSetWithdrawQueue is a free log subscription operation binding the contract event 0xe0c2db6b54586be6d7d49943139fccf0dd315ba63e55364a76c73cd8fdba724d.
//
// Solidity: event SetWithdrawQueue(address indexed caller, bytes32[] newWithdrawQueue)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSetWithdrawQueue(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSetWithdrawQueue, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "SetWithdrawQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSetWithdrawQueue)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "SetWithdrawQueue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetWithdrawQueue is a log parse operation binding the contract event 0xe0c2db6b54586be6d7d49943139fccf0dd315ba63e55364a76c73cd8fdba724d.
//
// Solidity: event SetWithdrawQueue(address indexed caller, bytes32[] newWithdrawQueue)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSetWithdrawQueue(log types.Log) (*MetaMorphoVaultSetWithdrawQueue, error) {
	event := new(MetaMorphoVaultSetWithdrawQueue)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "SetWithdrawQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSkimIterator is returned from FilterSkim and is used to iterate over the raw logs and unpacked data for Skim events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSkimIterator struct {
	Event *MetaMorphoVaultSkim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSkimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSkim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSkim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSkimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSkimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSkim represents a Skim event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSkim struct {
	Caller common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSkim is a free log retrieval operation binding the contract event 0x2ae72b44f59d038340fca5739135a1d51fc5ab720bb02d983e4c5ff4119ca7b8.
//
// Solidity: event Skim(address indexed caller, address indexed token, uint256 amount)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSkim(opts *bind.FilterOpts, caller []common.Address, token []common.Address) (*MetaMorphoVaultSkimIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "Skim", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSkimIterator{contract: _MetaMorphoVault.contract, event: "Skim", logs: logs, sub: sub}, nil
}

// WatchSkim is a free log subscription operation binding the contract event 0x2ae72b44f59d038340fca5739135a1d51fc5ab720bb02d983e4c5ff4119ca7b8.
//
// Solidity: event Skim(address indexed caller, address indexed token, uint256 amount)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSkim(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSkim, caller []common.Address, token []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "Skim", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSkim)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "Skim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSkim is a log parse operation binding the contract event 0x2ae72b44f59d038340fca5739135a1d51fc5ab720bb02d983e4c5ff4119ca7b8.
//
// Solidity: event Skim(address indexed caller, address indexed token, uint256 amount)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSkim(log types.Log) (*MetaMorphoVaultSkim, error) {
	event := new(MetaMorphoVaultSkim)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "Skim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSubmitCapIterator is returned from FilterSubmitCap and is used to iterate over the raw logs and unpacked data for SubmitCap events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSubmitCapIterator struct {
	Event *MetaMorphoVaultSubmitCap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSubmitCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSubmitCap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSubmitCap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSubmitCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSubmitCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSubmitCap represents a SubmitCap event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSubmitCap struct {
	Caller common.Address
	Id     [32]byte
	Cap    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitCap is a free log retrieval operation binding the contract event 0xe851bb5856808a50efd748be463b8f35bcfb5ec74c5bfde776fe0a4d2a26db27.
//
// Solidity: event SubmitCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSubmitCap(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MetaMorphoVaultSubmitCapIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "SubmitCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSubmitCapIterator{contract: _MetaMorphoVault.contract, event: "SubmitCap", logs: logs, sub: sub}, nil
}

// WatchSubmitCap is a free log subscription operation binding the contract event 0xe851bb5856808a50efd748be463b8f35bcfb5ec74c5bfde776fe0a4d2a26db27.
//
// Solidity: event SubmitCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSubmitCap(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSubmitCap, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "SubmitCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSubmitCap)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "SubmitCap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmitCap is a log parse operation binding the contract event 0xe851bb5856808a50efd748be463b8f35bcfb5ec74c5bfde776fe0a4d2a26db27.
//
// Solidity: event SubmitCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSubmitCap(log types.Log) (*MetaMorphoVaultSubmitCap, error) {
	event := new(MetaMorphoVaultSubmitCap)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "SubmitCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSubmitGuardianIterator is returned from FilterSubmitGuardian and is used to iterate over the raw logs and unpacked data for SubmitGuardian events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSubmitGuardianIterator struct {
	Event *MetaMorphoVaultSubmitGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSubmitGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSubmitGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSubmitGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSubmitGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSubmitGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSubmitGuardian represents a SubmitGuardian event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSubmitGuardian struct {
	NewGuardian common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmitGuardian is a free log retrieval operation binding the contract event 0x7633313af54753bce8a149927263b1a55eba857ba4ef1d13c6aee25d384d3c4b.
//
// Solidity: event SubmitGuardian(address indexed newGuardian)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSubmitGuardian(opts *bind.FilterOpts, newGuardian []common.Address) (*MetaMorphoVaultSubmitGuardianIterator, error) {

	var newGuardianRule []interface{}
	for _, newGuardianItem := range newGuardian {
		newGuardianRule = append(newGuardianRule, newGuardianItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "SubmitGuardian", newGuardianRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSubmitGuardianIterator{contract: _MetaMorphoVault.contract, event: "SubmitGuardian", logs: logs, sub: sub}, nil
}

// WatchSubmitGuardian is a free log subscription operation binding the contract event 0x7633313af54753bce8a149927263b1a55eba857ba4ef1d13c6aee25d384d3c4b.
//
// Solidity: event SubmitGuardian(address indexed newGuardian)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSubmitGuardian(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSubmitGuardian, newGuardian []common.Address) (event.Subscription, error) {

	var newGuardianRule []interface{}
	for _, newGuardianItem := range newGuardian {
		newGuardianRule = append(newGuardianRule, newGuardianItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "SubmitGuardian", newGuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSubmitGuardian)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "SubmitGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmitGuardian is a log parse operation binding the contract event 0x7633313af54753bce8a149927263b1a55eba857ba4ef1d13c6aee25d384d3c4b.
//
// Solidity: event SubmitGuardian(address indexed newGuardian)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSubmitGuardian(log types.Log) (*MetaMorphoVaultSubmitGuardian, error) {
	event := new(MetaMorphoVaultSubmitGuardian)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "SubmitGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSubmitMarketRemovalIterator is returned from FilterSubmitMarketRemoval and is used to iterate over the raw logs and unpacked data for SubmitMarketRemoval events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSubmitMarketRemovalIterator struct {
	Event *MetaMorphoVaultSubmitMarketRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSubmitMarketRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSubmitMarketRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSubmitMarketRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSubmitMarketRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSubmitMarketRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSubmitMarketRemoval represents a SubmitMarketRemoval event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSubmitMarketRemoval struct {
	Caller common.Address
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitMarketRemoval is a free log retrieval operation binding the contract event 0x3240fc70754c5a2b4dab10bf7081a00024bfc8491581ee3d355360ec0dd91f16.
//
// Solidity: event SubmitMarketRemoval(address indexed caller, bytes32 indexed id)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSubmitMarketRemoval(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*MetaMorphoVaultSubmitMarketRemovalIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "SubmitMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSubmitMarketRemovalIterator{contract: _MetaMorphoVault.contract, event: "SubmitMarketRemoval", logs: logs, sub: sub}, nil
}

// WatchSubmitMarketRemoval is a free log subscription operation binding the contract event 0x3240fc70754c5a2b4dab10bf7081a00024bfc8491581ee3d355360ec0dd91f16.
//
// Solidity: event SubmitMarketRemoval(address indexed caller, bytes32 indexed id)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSubmitMarketRemoval(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSubmitMarketRemoval, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "SubmitMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSubmitMarketRemoval)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "SubmitMarketRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmitMarketRemoval is a log parse operation binding the contract event 0x3240fc70754c5a2b4dab10bf7081a00024bfc8491581ee3d355360ec0dd91f16.
//
// Solidity: event SubmitMarketRemoval(address indexed caller, bytes32 indexed id)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSubmitMarketRemoval(log types.Log) (*MetaMorphoVaultSubmitMarketRemoval, error) {
	event := new(MetaMorphoVaultSubmitMarketRemoval)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "SubmitMarketRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultSubmitTimelockIterator is returned from FilterSubmitTimelock and is used to iterate over the raw logs and unpacked data for SubmitTimelock events raised by the MetaMorphoVault contract.
type MetaMorphoVaultSubmitTimelockIterator struct {
	Event *MetaMorphoVaultSubmitTimelock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultSubmitTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultSubmitTimelock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultSubmitTimelock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultSubmitTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultSubmitTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultSubmitTimelock represents a SubmitTimelock event raised by the MetaMorphoVault contract.
type MetaMorphoVaultSubmitTimelock struct {
	NewTimelock *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmitTimelock is a free log retrieval operation binding the contract event 0xb3aa0ade2442acf51d06713c2d1a5a3ec0373cce969d42b53f4689f97bccf380.
//
// Solidity: event SubmitTimelock(uint256 newTimelock)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterSubmitTimelock(opts *bind.FilterOpts) (*MetaMorphoVaultSubmitTimelockIterator, error) {

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "SubmitTimelock")
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultSubmitTimelockIterator{contract: _MetaMorphoVault.contract, event: "SubmitTimelock", logs: logs, sub: sub}, nil
}

// WatchSubmitTimelock is a free log subscription operation binding the contract event 0xb3aa0ade2442acf51d06713c2d1a5a3ec0373cce969d42b53f4689f97bccf380.
//
// Solidity: event SubmitTimelock(uint256 newTimelock)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchSubmitTimelock(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultSubmitTimelock) (event.Subscription, error) {

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "SubmitTimelock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultSubmitTimelock)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "SubmitTimelock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmitTimelock is a log parse operation binding the contract event 0xb3aa0ade2442acf51d06713c2d1a5a3ec0373cce969d42b53f4689f97bccf380.
//
// Solidity: event SubmitTimelock(uint256 newTimelock)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseSubmitTimelock(log types.Log) (*MetaMorphoVaultSubmitTimelock, error) {
	event := new(MetaMorphoVaultSubmitTimelock)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "SubmitTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MetaMorphoVault contract.
type MetaMorphoVaultTransferIterator struct {
	Event *MetaMorphoVaultTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultTransfer represents a Transfer event raised by the MetaMorphoVault contract.
type MetaMorphoVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MetaMorphoVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultTransferIterator{contract: _MetaMorphoVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultTransfer)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseTransfer(log types.Log) (*MetaMorphoVaultTransfer, error) {
	event := new(MetaMorphoVaultTransfer)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultUpdateLastTotalAssetsIterator is returned from FilterUpdateLastTotalAssets and is used to iterate over the raw logs and unpacked data for UpdateLastTotalAssets events raised by the MetaMorphoVault contract.
type MetaMorphoVaultUpdateLastTotalAssetsIterator struct {
	Event *MetaMorphoVaultUpdateLastTotalAssets // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultUpdateLastTotalAssetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultUpdateLastTotalAssets)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultUpdateLastTotalAssets)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultUpdateLastTotalAssetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultUpdateLastTotalAssetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultUpdateLastTotalAssets represents a UpdateLastTotalAssets event raised by the MetaMorphoVault contract.
type MetaMorphoVaultUpdateLastTotalAssets struct {
	UpdatedTotalAssets *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateLastTotalAssets is a free log retrieval operation binding the contract event 0x15c027cc4fd826d986cad358803439f7326d3aa4ed969ff90dbee4bc150f68e9.
//
// Solidity: event UpdateLastTotalAssets(uint256 updatedTotalAssets)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterUpdateLastTotalAssets(opts *bind.FilterOpts) (*MetaMorphoVaultUpdateLastTotalAssetsIterator, error) {

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "UpdateLastTotalAssets")
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultUpdateLastTotalAssetsIterator{contract: _MetaMorphoVault.contract, event: "UpdateLastTotalAssets", logs: logs, sub: sub}, nil
}

// WatchUpdateLastTotalAssets is a free log subscription operation binding the contract event 0x15c027cc4fd826d986cad358803439f7326d3aa4ed969ff90dbee4bc150f68e9.
//
// Solidity: event UpdateLastTotalAssets(uint256 updatedTotalAssets)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchUpdateLastTotalAssets(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultUpdateLastTotalAssets) (event.Subscription, error) {

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "UpdateLastTotalAssets")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultUpdateLastTotalAssets)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "UpdateLastTotalAssets", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateLastTotalAssets is a log parse operation binding the contract event 0x15c027cc4fd826d986cad358803439f7326d3aa4ed969ff90dbee4bc150f68e9.
//
// Solidity: event UpdateLastTotalAssets(uint256 updatedTotalAssets)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseUpdateLastTotalAssets(log types.Log) (*MetaMorphoVaultUpdateLastTotalAssets, error) {
	event := new(MetaMorphoVaultUpdateLastTotalAssets)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "UpdateLastTotalAssets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMorphoVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the MetaMorphoVault contract.
type MetaMorphoVaultWithdrawIterator struct {
	Event *MetaMorphoVaultWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMorphoVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMorphoVaultWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMorphoVaultWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMorphoVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMorphoVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMorphoVaultWithdraw represents a Withdraw event raised by the MetaMorphoVault contract.
type MetaMorphoVaultWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*MetaMorphoVaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MetaMorphoVaultWithdrawIterator{contract: _MetaMorphoVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MetaMorphoVaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MetaMorphoVault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMorphoVaultWithdraw)
				if err := _MetaMorphoVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_MetaMorphoVault *MetaMorphoVaultFilterer) ParseWithdraw(log types.Log) (*MetaMorphoVaultWithdraw, error) {
	event := new(MetaMorphoVaultWithdraw)
	if err := _MetaMorphoVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
