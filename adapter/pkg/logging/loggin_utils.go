/*
 *  Copyright (c) 2023, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package logging

import (
	"fmt"
)

// GetErrorByCode used to keep error details for error logs
func GetErrorByCode(code int, args ...interface{}) ErrorDetails {
	errorLog, ok := Mapper[code]
	if !ok {
		errorLog = ErrorDetails{
			ErrorCode: 0000,
			Message:   fmt.Sprintf("No error message found for error code: %v", code),
			Severity:  "MINOR",
		}
	}
	errorLog.Message = fmt.Sprintf(errorLog.Message, args...)
	return errorLog
}
