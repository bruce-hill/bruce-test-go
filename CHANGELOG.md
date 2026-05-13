# Changelog

## 2.1.0 (2026-05-13)

Full Changelog: [v2.0.0...v2.1.0](https://github.com/bruce-hill/bruce-test-go/compare/v2.0.0...v2.1.0)

### Features

* **internal:** support comma format in multipart form encoding ([5735c05](https://github.com/bruce-hill/bruce-test-go/commit/5735c05b6acc851e5be5b1c6cea0665adeb3ef4e))
* support setting headers via env ([3dbf3b9](https://github.com/bruce-hill/bruce-test-go/commit/3dbf3b968e9873d367cd8f97f91cce8f702a4c62))


### Bug Fixes

* fix issue with unmarshaling in some cases ([b289dcf](https://github.com/bruce-hill/bruce-test-go/commit/b289dcfbfe1b0b281116d3b1667d6d4917c8a8d0))
* fixes for pagination and iteration, plus iter.Seq support ([5e46895](https://github.com/bruce-hill/bruce-test-go/commit/5e46895d8f0ad3e90f38bea6526e107249a9051a))
* **go:** avoid panic when http.DefaultTransport is wrapped ([f9f5602](https://github.com/bruce-hill/bruce-test-go/commit/f9f560224e4a9dd01e1874a9d81295d35ba5ecac))
* prevent duplicate ? in query params ([6d7cd46](https://github.com/bruce-hill/bruce-test-go/commit/6d7cd46e9b05e6d8e7ffabbf9bf82712ea158452))


### Chores

* avoid embedding reflect.Type for dead code elimination ([aa5d8d6](https://github.com/bruce-hill/bruce-test-go/commit/aa5d8d6cbe92a87960262d6adf1e940821e05c05))
* **ci:** skip lint on metadata-only changes ([ed132d3](https://github.com/bruce-hill/bruce-test-go/commit/ed132d3bbec1c49d39946a059df2be38a5d716ab))
* **ci:** support opting out of skipping builds on metadata-only commits ([4eb92b1](https://github.com/bruce-hill/bruce-test-go/commit/4eb92b195f8ac0d9a39eca11fce1949463f104f1))
* **client:** fix multipart serialisation of Default() fields ([baebe8c](https://github.com/bruce-hill/bruce-test-go/commit/baebe8cb52e72efccd4dc94e676238bb6aed19a9))
* **internal:** codegen related update ([4b423f2](https://github.com/bruce-hill/bruce-test-go/commit/4b423f20b5619410aa7ddb7d350e6ff5d38c2567))
* **internal:** codegen related update ([c57f1f2](https://github.com/bruce-hill/bruce-test-go/commit/c57f1f234471f387ca3d1ac891c1b9244535640a))
* **internal:** more robust bootstrap script ([ce753ff](https://github.com/bruce-hill/bruce-test-go/commit/ce753ffb1b0921b9c4daadc6bb7a0b8412647c5a))
* **internal:** support default value struct tag ([3d38e01](https://github.com/bruce-hill/bruce-test-go/commit/3d38e0145d63c8e240e178a804a25a36316e7f8e))
* **internal:** tweak CI branches ([f1d5ee8](https://github.com/bruce-hill/bruce-test-go/commit/f1d5ee844a4d6a8ec4e3a23a670fec4631aa3de9))
* **internal:** update gitignore ([3c64e59](https://github.com/bruce-hill/bruce-test-go/commit/3c64e593110ac82c9a3a81bf022ac0d39b55a547))
* redact api-key headers in debug logs ([6422f59](https://github.com/bruce-hill/bruce-test-go/commit/6422f59e4cd789a5bfd0165a11bab0a0c8fe737d))
* remove unnecessary error check for url parsing ([fe42ac4](https://github.com/bruce-hill/bruce-test-go/commit/fe42ac403f52c8b1e293d580ccc23a3c45199b4a))
* **tests:** bump steady to v0.19.4 ([d864ea4](https://github.com/bruce-hill/bruce-test-go/commit/d864ea479a4ab662110601f2fe518f48d4fe4f69))
* **tests:** bump steady to v0.19.5 ([748e482](https://github.com/bruce-hill/bruce-test-go/commit/748e482f29cc9d15992a12b62a22b2488e24c4ea))
* **tests:** bump steady to v0.19.6 ([6246f76](https://github.com/bruce-hill/bruce-test-go/commit/6246f765c05717ce68339c7480302a6d0be1e6a6))
* **tests:** bump steady to v0.20.1 ([9799e59](https://github.com/bruce-hill/bruce-test-go/commit/9799e59d00f46a2d10e5e8ce76eb08657a4ca6fa))
* **tests:** bump steady to v0.20.2 ([d755ad4](https://github.com/bruce-hill/bruce-test-go/commit/d755ad4e2a8f7ca9af13944f30a8408cee964b0e))
* **tests:** bump steady to v0.22.1 ([7777093](https://github.com/bruce-hill/bruce-test-go/commit/7777093b531ea3d2fdfc00aafe1625d65669dda0))
* update docs for api:"required" ([1037894](https://github.com/bruce-hill/bruce-test-go/commit/1037894d624bf93d08627cc8800572a1e66b1f39))


### Refactors

* **tests:** switch from prism to steady ([fc08e5c](https://github.com/bruce-hill/bruce-test-go/commit/fc08e5c0718b5fb7c208faacd57301ea7f84e76b))

## 2.0.0 (2026-03-13)

Full Changelog: [v0.0.1...v2.0.0](https://github.com/bruce-hill/bruce-test-go/compare/v0.0.1...v2.0.0)

### Features

* **api:** manual updates ([a14d98e](https://github.com/bruce-hill/bruce-test-go/commit/a14d98eb77c38f956a2c5d92745f534eabeb8184))
* **api:** manual updates ([54321fe](https://github.com/bruce-hill/bruce-test-go/commit/54321fede1826e4d202233599e71c53723938a99))
* **api:** manual updates ([b5f1058](https://github.com/bruce-hill/bruce-test-go/commit/b5f1058f3be99aef4ebf258a70f5ebd819813674))


### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([2c087a6](https://github.com/bruce-hill/bruce-test-go/commit/2c087a68e6cba092a41d981e8d2ac9f9c1dbe4dc))


### Chores

* bump gjson version ([bb5deb6](https://github.com/bruce-hill/bruce-test-go/commit/bb5deb6f3ddd5a494f131ae8686d4bb674e983c6))
* **internal:** grammar fix (it's -&gt; its) ([81f6c3d](https://github.com/bruce-hill/bruce-test-go/commit/81f6c3daeb421fbd4bcc4da5432fb0e9f2f88ff9))
