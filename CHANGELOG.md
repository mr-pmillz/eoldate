# Changelog

All notable changes to this project will be documented [here](https://github.com/mr-pmillz/eoldate/blob/main/CHANGELOG.md?ref_type=heads)

## [1.1.4](https://github.com/mr-pmillz/eoldate/compare/v1.1.3...v1.1.4) - 2026-06-18

### ✨ New features

- Feat: add 60s timeout to eoldate HTTP client - ([d17cbc9](https://github.com/mr-pmillz/eoldate/commit/d17cbc93169265741855b02e7826881aa7cd84d1))

### 🛠 Improvements

- Update cliff.toml - ([2f084cc](https://github.com/mr-pmillz/eoldate/commit/2f084cc87cd0c8ccdc42a490bddd616925ba2fba))

### ⚙️  Miscellaneous

- Version bump - ([c795479](https://github.com/mr-pmillz/eoldate/commit/c795479867626dec56c0cd067c6aca7619a941a1))
- Sets http.Client.Timeout via a new HTTPTimeout constant so API requests to endoflife.date fail after 60s instead of hanging indefinitely. - ([d17cbc9](https://github.com/mr-pmillz/eoldate/commit/d17cbc93169265741855b02e7826881aa7cd84d1))

## [1.1.3](https://github.com/mr-pmillz/eoldate/compare/v1.1.2...v1.1.3) - 2026-06-18

### ⚙️  Miscellaneous

- Version bump - ([a71c64f](https://github.com/mr-pmillz/eoldate/commit/a71c64f1f2147d411c0205578fac93dcf0c6b0b8))
- Bump go version and pin action versions - ([f3c9ad1](https://github.com/mr-pmillz/eoldate/commit/f3c9ad1f3e507897dfec2ff6801f5ab00deadd37))

## [1.1.2](https://github.com/mr-pmillz/eoldate/compare/v1.1.1...v1.1.2) - 2026-06-18

### ⚙️  Miscellaneous

- Version bump - ([5090817](https://github.com/mr-pmillz/eoldate/commit/5090817ed54bd97bfcb99dc6a917fa967b4448f1))
- Replaced tablewriter with pterm for table output - ([a622934](https://github.com/mr-pmillz/eoldate/commit/a6229341792121923d87fb9d6ef7e7f12f154ce8))
- Update changelog - ([ef75c3f](https://github.com/mr-pmillz/eoldate/commit/ef75c3f6d84a7fa8d9d69ad8814f4e522ecd42f3))

## [1.1.1](https://github.com/mr-pmillz/eoldate/compare/v1.1.0...v1.1.1) - 2025-09-26

### ✨ New features

- Added missing entrypoint.sh - ([9f42407](https://github.com/mr-pmillz/eoldate/commit/9f42407cfa198ec2265a289d557e2b0b2cbd8b54))

## [1.1.0](https://github.com/mr-pmillz/eoldate/compare/v1.0.9...v1.1.0) - 2025-09-26

### 🛠 Improvements

- Update ci fixed build path - ([92fd2ba](https://github.com/mr-pmillz/eoldate/commit/92fd2bad7df6ae83d9a54e5006acb93b6c7a3050))

## [1.0.9](https://github.com/mr-pmillz/eoldate/compare/v1.0.8...v1.0.9) - 2025-09-26

### 🛠 Improvements

- Update ci and version bump again - ([f334992](https://github.com/mr-pmillz/eoldate/commit/f334992e04fecc84dd851b590a09da1cb6fd2f97))

## [1.0.8](https://github.com/mr-pmillz/eoldate/compare/v1.0.7...v1.0.8) - 2025-09-26

### 🛠 Improvements

- Update ci and version bump - ([92310aa](https://github.com/mr-pmillz/eoldate/commit/92310aa95b09d4d316f17dfdf0a4cce21cd8327f))

## [1.0.7](https://github.com/mr-pmillz/eoldate/compare/v1.0.6...v1.0.7) - 2025-09-26

### 🛠 Improvements

- Update MinJavaVersion type - ([85f7baf](https://github.com/mr-pmillz/eoldate/commit/85f7baffc00a7110bae1d0628c21605bb452610a))
- Update changelog - ([b252776](https://github.com/mr-pmillz/eoldate/commit/b2527760419bef80ef8f403b532e27e53c263ead))

## [1.0.6](https://github.com/mr-pmillz/eoldate/compare/v1.0.5...v1.0.6) - 2024-09-29

### 🛠 Improvements

- Updated IsSupportedSoftwareVersion to return a *semver.Version object - ([1e8c7dd](https://github.com/mr-pmillz/eoldate/commit/1e8c7dd0529d8a73742a63cca117dc1e5274d8f1))
- Update changelog - ([77c28e9](https://github.com/mr-pmillz/eoldate/commit/77c28e9f37d3965aa2d4568f282392b94209f415))

### ⚙️  Miscellaneous

- Dep bumped goreleaser to v2 in ci - ([1e8c7dd](https://github.com/mr-pmillz/eoldate/commit/1e8c7dd0529d8a73742a63cca117dc1e5274d8f1))

## [1.0.5](https://github.com/mr-pmillz/eoldate/compare/v1.0.4...v1.0.5) - 2024-09-28

### ✨ New features

- Add color to extended support date - ([37d435a](https://github.com/mr-pmillz/eoldate/commit/37d435affba8517c65518d355a18fe4a53347cff))

### 🛠 Improvements

- Update CHANGELOG.md - ([2e29850](https://github.com/mr-pmillz/eoldate/commit/2e298500111b603722f97ae66124dd045c99fbe9))

## [1.0.4](https://github.com/mr-pmillz/eoldate/compare/v1.0.3...v1.0.4) - 2024-09-11

### 🛠 Improvements

- Update CHANGELOG.md - ([c0c4310](https://github.com/mr-pmillz/eoldate/commit/c0c4310fbff5d2320d69437d838058a72dd2b14b))

### ⚙️  Miscellaneous

- Return matched Product data to function calls - ([ec4ee11](https://github.com/mr-pmillz/eoldate/commit/ec4ee11b9a4bd34b4ba2e86d89f05bcefbc04a1f))

## [1.0.3](https://github.com/mr-pmillz/eoldate/compare/v1.0.2...v1.0.3) - 2024-09-10

### ⚡ Performance

- Optimized latest version edge case handling - ([9fef918](https://github.com/mr-pmillz/eoldate/commit/9fef91897a92bea3265e3dabd5efd79d839c54d8))

### 📦 Dependencies

- Dependency bumped versions - ([9fef918](https://github.com/mr-pmillz/eoldate/commit/9fef91897a92bea3265e3dabd5efd79d839c54d8))

### 🛠 Improvements

- Update CHANGELOG.md - ([3ff3992](https://github.com/mr-pmillz/eoldate/commit/3ff3992791c68475b2cc62e971987bd30bb507f0))

### ⚙️  Miscellaneous

- Removed comment from README.md example code - ([6fa085a](https://github.com/mr-pmillz/eoldate/commit/6fa085a8275b4a4f1065bbf0f4d0434681d6aaed))

## [1.0.2](https://github.com/mr-pmillz/eoldate/compare/v1.0.1...v1.0.2) - 2024-09-10

### ✨ New features

- Added support for versions less than lowest cycle version - ([3b0125d](https://github.com/mr-pmillz/eoldate/commit/3b0125ddc7014ad0d20d2ce8b1975eecf3357b23))

### 🛠 Improvements

- Updated README.md example code - ([995b24e](https://github.com/mr-pmillz/eoldate/commit/995b24efc0a182ed3e320ebaf226581a80974f8b))
- Update CHANGELOG.md - ([5e79a8a](https://github.com/mr-pmillz/eoldate/commit/5e79a8ab26e150d3f7483cba4579656adfa60db5))

## [1.0.1](https://github.com/mr-pmillz/eoldate/compare/v1.0.0...v1.0.1) - 2024-09-10

### ✨ New features

- Added semver for semantic versioning support compairsons - ([5aca6c4](https://github.com/mr-pmillz/eoldate/commit/5aca6c4a4cffcebb683dbefc1bdb472cb33a0b3d))

### ⚡ Performance

- Optimized IsSupportedSoftwareVersion and IsVersionSupported funcs - ([5aca6c4](https://github.com/mr-pmillz/eoldate/commit/5aca6c4a4cffcebb683dbefc1bdb472cb33a0b3d))

### 🛠 Improvements

- Update CHANGELOG.md - ([c229b83](https://github.com/mr-pmillz/eoldate/commit/c229b83431eba0723561ae4456e28ee5dbafeab5))

## [1.0.0](https://github.com/mr-pmillz/eoldate/compare/v0.0.9...v1.0.0) - 2024-09-10

### ✨ New features

- Added local file caching - ([248e77c](https://github.com/mr-pmillz/eoldate/commit/248e77cba1203a48dd4033c1ca46334721e2ec34))

### 🛠 Improvements

- Update CHANGELOG.md - ([0416601](https://github.com/mr-pmillz/eoldate/commit/04166019743f59748817824f60e38e8df5d83eb9))

## [0.0.9](https://github.com/mr-pmillz/eoldate/compare/v0.0.8...v0.0.9) - 2024-09-05

### 🛠 Improvements

- Update CHANGELOG.md - ([2bf0d92](https://github.com/mr-pmillz/eoldate/commit/2bf0d922f0f6fde30e4d5435dac4760f85213a35))

### ⚙️  Miscellaneous

- Version bump - ([ed8760a](https://github.com/mr-pmillz/eoldate/commit/ed8760a8f2e434807a8ebdcaf182e7253e9cbeff))
- Resolved linter warnings - ([1264fb9](https://github.com/mr-pmillz/eoldate/commit/1264fb9ac842226f201191ccf7fe241dfc5dec5f))

## [0.0.8](https://github.com/mr-pmillz/eoldate/compare/v0.0.7...v0.0.8) - 2024-09-05

### 🛠 Improvements

- Update CHANGELOG.md - ([45feb89](https://github.com/mr-pmillz/eoldate/commit/45feb8993a3e6971909b1d48c116c2e79ea57386))

### ⚙️  Miscellaneous

- Version bump - ([81c5229](https://github.com/mr-pmillz/eoldate/commit/81c52299825e9c39e3ded6a0a0c56a8c47aa37ab))
- Extra crispy crucial :spy: - ([3c47225](https://github.com/mr-pmillz/eoldate/commit/3c47225cd5ac736b8f74a3e6de5dffec3d8b61be))

## [0.0.7](https://github.com/mr-pmillz/eoldate/compare/v0.0.6...v0.0.7) - 2024-09-05

### 🛠 Improvements

- Update CHANGELOG.md - ([e1594de](https://github.com/mr-pmillz/eoldate/commit/e1594deb64b02903892a87140daaf8ebb4dcd8b0))

### ⚙️  Miscellaneous

- Much moe betta - ([4d38a27](https://github.com/mr-pmillz/eoldate/commit/4d38a27a87471e14ae25002146de498d76b2ac7e))

## [0.0.6](https://github.com/mr-pmillz/eoldate/compare/v0.0.5...v0.0.6) - 2024-09-05

### ⚡ Performance

- Optimized code, dynamic table columns - ([55f3590](https://github.com/mr-pmillz/eoldate/commit/55f3590510073e402c9ba8ab234c13d56b082535))

### 🛠 Improvements

- Update README.md CHANGELOG.md cliff.toml - ([d398e29](https://github.com/mr-pmillz/eoldate/commit/d398e29f3840565fbf14dc86841049681a668882))
- Update CHANGELOG.md - ([0d148dd](https://github.com/mr-pmillz/eoldate/commit/0d148dd3f92802ec48cd1e4933dc9376d83112e8))

### ⚙️  Miscellaneous

- Moe betta - ([b315a3e](https://github.com/mr-pmillz/eoldate/commit/b315a3e7ce6db0c9ef4b7f27e13d994e2f4947d6))

## [0.0.5](https://github.com/mr-pmillz/eoldate/compare/v0.0.4...v0.0.5) - 2024-09-04

### ✨ New features

- Added additional product types, link, extended support - ([2bfc22d](https://github.com/mr-pmillz/eoldate/commit/2bfc22df5bc2d6299a4bcce803ace4547ffec8d3))

## [0.0.4] - 2024-09-03

### 🧪 Testing

- Testing github goreleaser release notes from git-cliff artifact - ([efbb3e6](https://github.com/mr-pmillz/eoldate/commit/efbb3e60ec3e9a4da09ec06a8d8c3948c21d1ed8))

### 🛠 Improvements

- Update CHANGELOG.md - ([3dfaae1](https://github.com/mr-pmillz/eoldate/commit/3dfaae17b22d03532f36613d2d2174085f34d6dc))

## [0.0.3](https://github.com/mr-pmillz/eoldate/compare/v0.0.2...v0.0.3) - 2024-09-03

### ✨ New features

- Added changelog generation to cicd - ([bdc60ab](https://github.com/mr-pmillz/eoldate/commit/bdc60ab253fd8bbeb6d3cce27c94d7ac91bd12ec))

### ⚙️  Miscellaneous

- Show usage when no flags specified - ([8bf2224](https://github.com/mr-pmillz/eoldate/commit/8bf22248442e7e5b45d9bd2987bf675201dc3d77))

## [0.0.2](https://github.com/mr-pmillz/eoldate/compare/v0.0.1...v0.0.2) - 2024-09-03

### ⚙️  Miscellaneous

- Release v0.0.2 - ([855ff22](https://github.com/mr-pmillz/eoldate/commit/855ff2272fa8ae94493965630314f9c047c64149))

## [0.0.1] - 2024-09-03

### ⚙️  Miscellaneous

- Initial poc - ([874857a](https://github.com/mr-pmillz/eoldate/commit/874857a0c84039ef284ccb9d7b7f82ca6967be4e))
- Initial commit - ([9a0c65c](https://github.com/mr-pmillz/eoldate/commit/9a0c65c4587c942f4013fd1a54b1925770c235f6))

<!-- generated by git-cliff -->
