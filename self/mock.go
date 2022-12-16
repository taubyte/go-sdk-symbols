//go:build !wasi && !wasm
// +build !wasi,!wasm

package selfSym

import (
	"github.com/taubyte/go-sdk-symbols/mocks"
)

func MockProject(testProjectId string) {
	ProjectSize = mocks.SizeStringFunction(testProjectId)
	Project = mocks.DataStringFunction(testProjectId)
}

func MockApplication(testApplicationId string) {
	ApplicationSize = mocks.SizeStringFunction(testApplicationId)
	Application = mocks.DataStringFunction(testApplicationId)
}

func MockFunction(testFunctionId string) {
	IdSize = mocks.SizeStringFunction(testFunctionId)
	Id = mocks.DataStringFunction(testFunctionId)
}

func MockCommit(testCommitId string) {
	CommitSize = mocks.SizeStringFunction(testCommitId)
	Commit = mocks.DataStringFunction(testCommitId)
}

func MockBranch(testBranch string) {
	BranchSize = mocks.SizeStringFunction(testBranch)
	Branch = mocks.DataStringFunction(testBranch)
}
