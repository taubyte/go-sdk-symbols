//go:build !wasi && !wasm
// +build !wasi,!wasm

package ethereumSym

import (
	"crypto/rand"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/ethereum/client/bytes"
	ethCodec "github.com/taubyte/go-sdk/ethereum/client/codec"
	"github.com/taubyte/go-sdk/utils/codec"
)

func MockNewBoundContract(testContract MockContract, contractId uint32, inputWriteFailure, outputWriteFailure bool) error {
	encoded, err := mockNewContractHelper(testContract, inputWriteFailure, outputWriteFailure)
	if err != nil {
		return fmt.Errorf("Mock new contract helper failed with: %s", err)
	}

	EthNewContractSize = func(clientId uint32, abiPtr *byte, abiSize uint32, address string, methodsSizePtr, eventsSizePtr, contractPtr *uint32) (error errno.Error) {
		if clientId != testContract.ContractSizeClientId {
			return 1
		}

		*methodsSizePtr = uint32(len(encoded))
		*contractPtr = contractId
		return 0
	}

	return nil
}

func MockDeployContract(testContract MockContract, address string, transactionId, contractId uint32, inputWriteFailure, outputWriteFailure bool) error {
	methodsBytes, err := mockNewContractHelper(testContract, inputWriteFailure, outputWriteFailure)
	if err != nil {
		return fmt.Errorf("setting NewContract method to mock failed with: %s", err)
	}

	_address, err := bytes.AddressFromHex(address)
	if err != nil {
		return fmt.Errorf("getting address bytes failed with: %s", err)
	}

	EthDeployContract = func(clientId uint32, chainIdPtr *byte, chainIdSize uint32, bin string, abiPtr *byte, abiSize uint32, privKeyPtr *byte, privKeySize uint32, addressPtr *byte, methodsSizePtr, eventSizePtr, contractIdPtr, transactionIdPtr *uint32) (error errno.Error) {
		if clientId != testContract.ContractSizeClientId {
			return 1
		}

		data := unsafe.Slice(addressPtr, len(_address))
		copy(data, _address)
		*transactionIdPtr = transactionId
		*contractIdPtr = contractId
		*methodsSizePtr = uint32(len(methodsBytes))

		return 0
	}

	return nil
}

func MockTransactContract(testClientId uint32, transactionId uint32) {
	EthTransactContract = func(clientId, contractId uint32, chainIdPtr *byte, chainIdSize uint32, method string, privKeyPtr *byte, privKeySize uint32, inputPtr *byte, inputSize, isJson uint32, transactionIdPtr *uint32) (error errno.Error) {
		if clientId != testClientId {
			return 1
		}

		*transactionIdPtr = transactionId

		return 0
	}
}

func mockGetCallOutput(contract MockContract, method string, outputTypeFailure, outputLengthFailure bool) ([]byte, error) {
	var outputs [][]byte
	if outputLengthFailure {
		outputs = [][]byte{[]byte("incompatible"), []byte("extra output")}
	} else {
		if outputTypeFailure {
			outputs = [][]byte{[]byte("incompatible")}
		} else {
			for _, output := range contract.Methods[method].Outputs {
				encoder, err := ethCodec.Converter(reflect.TypeOf(output).String()).Encoder()
				if err != nil {
					return nil, fmt.Errorf("getting Converter for output %v failed with: %s", output, err)
				}

				data, err := encoder(output)
				if err != nil {
					return nil, fmt.Errorf("encoding data failed with: %s", err)
				}

				outputs = append(outputs, data)
			}
		}
	}

	var output []byte
	err := codec.Convert(outputs).To(&output)
	if err != nil {
		return nil, fmt.Errorf("converting output list to []byte failed with: %s", err)
	}

	return output, nil
}

func MockCall(contract MockContract, outputTypeFailure, outputLengthFailure bool) {
	EthCallContractSize = func(clientId, contractId uint32, method string, inputsPtr *byte, inputsSize, isJSON uint32, outputSizePtr *uint32) (error errno.Error) {
		if clientId != contract.CallSizeClientId {
			return 1
		}

		output, err := mockGetCallOutput(contract, method, outputTypeFailure, outputLengthFailure)
		if err != nil {
			return 1
		}

		*outputSizePtr = uint32(len(output))
		return 0
	}

	EthCallContract = func(clientId, contractId uint32, method string, outputPtr *byte) (error errno.Error) {
		if clientId != contract.CallDataClientId {
			return 1
		}

		output, err := mockGetCallOutput(contract, method, outputTypeFailure, outputLengthFailure)
		if err != nil {
			return 1
		}

		data := unsafe.Slice(outputPtr, len(output))
		copy(data, output)

		return 0
	}
}

type MockContract struct {
	Methods              map[string]MockContractMethod
	ContractSizeClientId uint32
	ContractDataClientId uint32
	MethodSizeClientId   uint32
	MethodDataClientId   uint32
	CallSizeClientId     uint32
	CallDataClientId     uint32
}

type MockContractMethod struct {
	Inputs  []interface{}
	Outputs []interface{}
}

func mockMethodInputOutputHelper(method string, contract MockContract, inputWriteFailure, outputWriteFailure bool) (inputs []byte, outputs []byte, err errno.Error) {
	if inputWriteFailure {
		inputs = make([]byte, 8)
		rand.Read(inputs)
	} else {
		var inputList []string
		for _, input := range contract.Methods[method].Inputs {
			inputList = append(inputList, reflect.TypeOf(input).String())
		}

		if err := codec.Convert(inputList).To(&inputs); err != nil {
			return nil, nil, 1
		}
	}

	if outputWriteFailure {
		outputs = make([]byte, 8)
		rand.Read(outputs)
	} else {
		var outputList []string
		for _, output := range contract.Methods[method].Outputs {
			outputList = append(outputList, reflect.TypeOf(output).String())
		}

		if err := codec.Convert(outputList).To(&outputs); err != nil {
			return nil, nil, 1
		}
	}

	return
}

func mockNewContractHelper(contract MockContract, inputWriteFailure, outputWriteFailure bool) ([]byte, error) {
	var encodedMethods []byte
	var methodNames []string

	for methodName := range contract.Methods {
		methodNames = append(methodNames, methodName)
	}

	if err := codec.Convert(methodNames).To(&encodedMethods); err != nil {
		return nil, err
	}

	EthNewContract = func(clientId, contractId uint32, methodsPtr, eventsPtr *byte) (error errno.Error) {
		if clientId != contract.ContractDataClientId {
			return 1
		}

		data := unsafe.Slice(methodsPtr, len(encodedMethods))
		copy(data, encodedMethods)
		return 0
	}

	EthGetContractMethodSize = func(clientId, contractId uint32, method string, inputSizePtr, outputSizePtr *uint32) (error errno.Error) {
		if clientId != contract.MethodSizeClientId {
			return 1
		}

		inputs, outputs, err := mockMethodInputOutputHelper(method, contract, inputWriteFailure, outputWriteFailure)
		if err != 0 {
			return err
		}

		*inputSizePtr = uint32(len(inputs))
		*outputSizePtr = uint32(len(outputs))

		return 0
	}

	EthGetContractMethod = func(clientId, contractId uint32, method string, inputPtr, outputPtr *byte) (error errno.Error) {
		if clientId != contract.MethodDataClientId {
			return 1
		}

		inputs, outputs, err := mockMethodInputOutputHelper(method, contract, inputWriteFailure, outputWriteFailure)
		if err != 0 {
			return err
		}

		if inputWriteFailure {
			d := unsafe.Slice(inputPtr, 12)
			copy(d, []byte("hello worlds"))
		} else {
			data := unsafe.Slice(inputPtr, len(inputs))
			copy(data, inputs)
		}

		if outputWriteFailure {
			d := unsafe.Slice(outputPtr, 12)
			copy(d, []byte("hello worlds"))
		} else {
			data := unsafe.Slice(outputPtr, len(outputs))
			copy(data, outputs)
		}

		return 0
	}

	return encodedMethods, nil
}
