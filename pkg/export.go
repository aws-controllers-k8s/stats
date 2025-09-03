// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package extractor

import (
	"encoding/json"
	"fmt"
	"os"
)

// WriteServiceOperationsJSON writes service operations to a JSON file
func WriteServiceOperationsJSON(serviceOps *ServiceOperations, outputPath string) error {
	data, err := json.MarshalIndent(serviceOps, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	
	return os.WriteFile(outputPath, data, 0644)
}