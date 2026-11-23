/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package file

import (
	"crypto/sha1" // #nosec
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/klog/v2"
)

// SHA1 returns the SHA1 of a file.
func SHA1(filename string) string {
	safeFilename, err := managedPath(filename)
	if err != nil {
		klog.ErrorS(err, "Refusing to hash unmanaged file path", "path", filename)
		return ""
	}

	// codeql[go/path-injection] managedPath restricts this path to the fixed SSL/auth directories.
	s, err := os.ReadFile(safeFilename)
	if err != nil {
		klog.ErrorS(err, "Error reading file", "path", filename)
		return ""
	}

	return SHA1Bytes(s)
}

// SHA1Bytes returns the SHA1 of data.
func SHA1Bytes(data []byte) string {
	hasher := sha1.New() // #nosec

	if _, err := hasher.Write(data); err != nil {
		return ""
	}

	return hex.EncodeToString(hasher.Sum(nil))
}

func managedPath(filename string) (string, error) {
	absFilename, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}

	for _, directory := range directories {
		base, err := filepath.Abs(directory)
		if err != nil {
			continue
		}

		relative, err := filepath.Rel(base, absFilename)
		if err == nil && relative != ".." && !strings.HasPrefix(relative, ".."+string(filepath.Separator)) && !filepath.IsAbs(relative) {
			return filepath.Join(base, filepath.Base(absFilename)), nil
		}
	}

	return "", fmt.Errorf("path %q is outside managed directories", filename)
}
