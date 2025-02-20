import { Assertion } from "chai";
import { ContractFunctionExecutionError } from "viem";

Assertion.addMethod(
  "revertedWithCustomError",
  async function (
    this: Chai.AssertionStatic,
    contract: any,
    expectedError: string
  ) {
    const promise = this._obj as Promise<unknown>;

    let errorCaught: ContractFunctionExecutionError | null = null;

    try {
      await promise;
    } catch (error) {
      if (error instanceof ContractFunctionExecutionError) {
        errorCaught = error;
      } else {
        throw error;
      }
    }

    this.assert(
      errorCaught !== null,
      `Expected transaction to revert with custom error '${expectedError}', but it did not revert.`,
      `Expected transaction *not* to revert with custom error '${expectedError}', but it did.`
    );

    this.assert(
      errorCaught?.details.includes(`custom error '${expectedError}()'`) ||
        errorCaught?.details.includes(`custom error '${expectedError}'`),
      `Expected revert detail to include "custom error '${expectedError}'" but got:\n${errorCaught?.details}`,
      `Expected revert detail *not* to include "custom error '${expectedError}'" but got:\n${errorCaught?.details}`
    );
  }
);
