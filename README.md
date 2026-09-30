# 红果鉴 / 真果鉴

> ⚠️ **免责声明**：本项目源码来自网上大名鼎鼎的**鱼佬**（原作者）。我只是把它拿来打包、测试着玩，方便自己用，**不保证任何可用性，随时可能删库**。

Flutter 多端独立短剧 / 影视应用，原名「短剧库 APP」。站源请求、解析、下载和播放均在设备上完成，不依赖自建服务。当前源码版本：**0.2.68+75**（开发快照，错播身份校验、媒体代理定向回归与全站源 iOS 云构建通过，本版已确认真机安装，错播与闲置恢复效果待用户复验）。

本轮在 0.2.50 基线上新增两个原生站源：**韩小圈**（`hanxiaoquan`，韩剧 / 韩国电影 / 韩国综艺 / 韩国动漫）与**鬼片网**（`guipian`，鬼片 / 电视剧 / 动漫）。默认可见站源顺序为：**红果 → 韩小圈 → 鬼片 → 青空**。

> 说明：应用长期按「只维护源码与定向测试，集中验证另行安排」的方式推进。未完成集中验证与真机验收的项目保持未验收状态；历史版本的检查记录不能作为后续新增功能的验收结论。

| 编译方式 | 应用名称 | 可用站源 |
| --- | --- | --- |
| 默认（不加参数） | 红果鉴 | 仅红果 |
| 加 `--all-sources` | 真果鉴 | 红果、韩小圈、鬼片、青空，以及黄豆、剧果、野果、帝果、黄果视频、黄果 AI、黄果旧版 |

默认可见站源（红果、韩小圈、鬼片、青空）无需密码；其余站源默认隐藏，在站源管理里通过密码锁解锁后显示。已有受限用户不会自动获得新增站源权限，管理员可在用户管理中勾选。

站源密码锁入口为连续点击底部“最近观看”6 次。首次可选择“不使用密码”显示全部站源；已设密码时，解锁后再次打开弹窗，可选择“关闭密码功能”。`0.2.64+71` 的 iOS 自动导出检查每 8 秒读取配置，导致临时解锁状态被清除，隐藏站源报用户权限错误并跳回红果；关闭密码功能可避开此问题。本地 `0.2.65+72` 已修改配置重读逻辑：同一用户、用户配置和站源密码未变化时保留本次运行的解锁状态，真正重启、密码变更或用户配置变化时仍恢复隐藏；已完成 `0.2.65+72` 全站源 iOS 云构建，真机修复效果待确认。

这是**编译选项**，应用内不能切换版本。标题、Android 桌面名称与电视横幅、Windows 窗口与分发文件名、iOS 显示名随编译选项变化。界面、站源调用、原生下载调度同时限制可用站源；红果版不会访问或继续执行其他站源的旧任务。

两版保留原 Android / iOS 应用标识及数据目录；Android 使用同一签名可相互覆盖升级，不能作为两个独立正式应用并排安装。切换版本保留追剧、观看记录、用户权限和下载记录；备份格式保持兼容。

## 使用

| 功能 | 操作 |
| --- | --- |
| 浏览 | 标题栏依次提供排序与筛选、榜单、多选下载和展开搜索。分类保持独立一行，内容区左右滑动切类；支持默认、名称、自然季号、上线日期、热度、播放量排序及连载状态筛选。列表下滑接近底部自动加载下一页，底部「加载更多」为兜底 |
| 搜索 | 红果合并官网与名称索引，输入停顿 300 毫秒显示最多 10 条联想；韩小圈、剧果、野果、鬼片调用各自的在线搜索并按分页继续加载。部分失败保留有效结果；最近 20 次搜索按用户保存 |
| 推荐 | 仅红果分类栏在「全部」后显示「推荐」，可切换真人剧、漫剧、AI 剧并分别记住位置 |
| 详情 | 普通点剧直接进入播放；下载和兜底入口保留详情页。详情页浏览封面、资料、追剧状态与可展开简介；底部固定「立即播放 / 继续播放」和下载入口。选集默认折叠，展开后每组 50 集，可切换范围、定位当前或输入集数跳转；宽屏与电视采用资料、选集分栏 |
| 追剧与历史 | 按想看 / 在看 / 已看 / 有更新筛选；红果已追剧发现新季后进入「有更新」。追剧和历史均可搜索。每个用户独立保存，每 5 秒及退出播放时记录真实进度 |
| 卡片与批量下载 | 卡片「更多」提供收藏、状态和下载选集；电脑右键，电视菜单键。发现页长按卡片或点标题栏「多选下载」，一次最多 50 部 |
| VIP | 全站源版浏览黄豆或全部站源时显示 VIP 图标，仅过滤黄豆；默认隐藏已确认的 VIP。未知、免费、VIP 分别保存，未知不会覆盖已知状态 |
| 手机播放 | 上滑下一集、下滑上一集；长按 350 毫秒临时 3 倍速。竖屏轻点播放 / 暂停，横屏轻点控制条，双击播放 / 暂停。视频下方为「选集 / 简介 / 下载」Tab |
| 画中画 | Android 手机 / 平板播放器接入画中画按钮；进入前隐藏自绘控制层与弹幕，小窗只保留画面。Android TV、Windows、iOS 暂未接入 |
| 画质增强 | 仅 Windows 桌面端显示并运行增强链路（关闭 / 自动 / 省电增强 / 清晰优先）；移动端与电视端隐藏 |
| 弹幕 | 红果在线播放默认开启；播放器侧边圆形「弹」字按钮可关闭、查看状态及失败重试。本地播放不加载弹幕 |
| 下载 | 播放页「下载」Tab 和详情页下载入口均可选择分集与画质；支持批量暂停、继续、重试、删除及清理任务但保留视频。「更新本剧」补新增或缺失分集 |
| 更多 | 站源管理、用户管理、设置与备份、界面模式、关于 |

| 站源 | 浏览与播放 | 搜索 |
| --- | --- | --- |
| 红果 | 真人剧、漫剧、AI 剧及分集 | 联网搜索与官网搜索联想 |
| 韩小圈 | MacCMS 模板：最新韩剧 / 韩国电影 / 韩国综艺 / 韩国动漫，多线路分集 | 站源在线搜索 |
| 鬼片 | MacCMS 站点：鬼片 / 电视剧 / 动漫，多线路分集 | RSS 最新条目标题匹配（站点搜索已停用） |
| 青空 | 番剧、剧场动画、特摄及分集 | 站源在线搜索 |
| 黄豆 | 列表、VIP 标记及分集 | 筛选已加载短剧 |
| 剧果 | 热门、最新、接口分类、详情及可播放分集；签名 Cookie 接入在线播放、预加载和下载 | 站源在线搜索，支持继续加载 |
| 野果 | 接口实际分类、目录、详情及真实分集；按分集重新取流，保留 H.264 / H.265 地址 | 站源在线分页搜索 |
| 帝果 | 网页分类、目录、详情及分集；vplayer 签名解析 | 站源在线分页搜索 |
| 黄果视频 / 黄果 AI / 黄果旧版 | 列表、分类、详情及分集 | 筛选已加载短剧 |

站源可用性、清晰度和区域限制取决于源站及网络；应用不解除源站 VIP 或其他授权限制。

### 网络与资源设置

2026-09-30 真机排查发现部分黄果 AI 剧集在 iPhone 直连时分集目录超时，开启手机代理后用户确认加载恢复且速度很快。遇到同类问题可先检查到该站源的网络连通性；“检测连接通过”与每部剧、每集都可播放并不等价。

管理员从「设置与备份 → 网络与资源」选择自动、直连或手动代理，手动地址支持 HTTP、HTTPS、SOCKS5 / SOCKS5h。代理地址默认遮蔽、不进入备份；设备内播放服务始终绕过代理。自动模式读取系统静态代理与排除列表，恢复前台及每 30 秒更新；代理不会关闭 TLS 校验。

目录请求并发可设 1–6、间隔 0–5000 毫秒（默认 3 / 250ms）；下载并发默认 2、可设 1–6。仍保留前台优先、后台让行、取消、退避及请求数量边界。

### 局域网互联：追剧同步与推送播放（已接入，待验证）

两台设备在同一局域网直连，自动同步追剧记录或接续播放，不依赖账号服务器。服务类型 `_zgj-link._tcp`（DNS-SD / mDNS），按设备 ID 合并多网卡地址，首次配对记录证书指纹，后续连接固定校验。

入口：「追剧 → 标题栏同步」或「设置与备份 → 设备互联」。手动同步默认「双向合并」，可选覆盖对方 / 覆盖本机，先「查看预览」再执行；自动同步仅发送改动条目与删除标记，播放期间约每 10 秒发送进度。冲突记录保留候选，用户从同步页逐项处理。

## 安装包与平台状态

| 平台 | 包与状态 |
| --- | --- |
| Android 8.0+ | 三架构（arm64-v8a / armeabi-v7a / x86_64）APK；同一签名可覆盖升级 |
| Windows 10/11 x64 | 完整 ZIP 解压后运行 `hongguojian.exe` / `zhenguojian.exe`，保留所有 DLL 与 `data`；局域网原生发现依赖 Windows 10 1903+ |
| Android TV | 与手机共用源码，自动识别电视模式并保持横屏；待电视 / 盒子实机验收 |
| iOS 15.1+ | 已加入工程、Go 核心链接、媒体依赖、文件管理与构建脚本；iOS 播放页禁用 media_kit_video 硬件纹理加速以规避 libmpv 渲染崩溃；2026-09-30 已完成 GitHub macOS 全站源版未签名 IPA 构建，用户已用自己的 Apple ID 自签安装并确认可启动，真机功能验收待完成 |

`INSTALL_FAILED_NO_MATCHING_ABIS` 表示 APK 与设备架构不匹配，请更换对应架构安装包。

### GitHub Actions

推送 `main` / `master`、`v*` 标签、提交 PR，或手动运行 **Build app packages**，会先检查再构建两版（默认与 `--all-sources`）：

| 产物 | 内容 |
| --- | --- |
| `*-android` | 三种架构 APK 和 SHA256 |
| `*-windows` | 完整 ZIP 和 SHA256 |
| `*-ios-unsigned` | 未签名 `.app` ZIP 和 SHA256，不能直接当已签名 IPA 安装 |

全站源 iOS 可单独手动运行 **Build iOS all sources**（`.github/workflows/ios-all-sources.yml`），先运行 `flutter pub get` 补齐上游新增依赖的锁定，再在 GitHub 的 macOS 环境执行 `python3 scripts/build_ios.py --all-sources`。构建成功后下载 `zhenguojian-ios-all-sources` 产物，其中包含未签名 IPA、`.app` ZIP、SHA256 和本次构建的 `pubspec.lock`；IPA 仍需使用自己的 Apple ID 自签后安装，产物保留 7 天。该任务不发布 Release。

已成功构建的开发快照包括 `0.2.64+71`（用户已自签安装并启动）和 `0.2.65+72`（包含密码锁刷新及 IPA 打包修复）：[0.2.65+72 GitHub Actions 构建记录](https://github.com/GaodYang/guoapp/actions/runs/36691024970)。`0.2.65+72` 下载后已核对云端 SHA256、`Payload/Runner.app` 结构及版本号。构建及 Go FFI 符号检查通过不代表真机功能验收；`--all-sources` 启用全部站源，不改变隐藏站源密码锁，也不补齐尚未实现的 iOS 画中画等平台功能。

`0.2.66+73` 修复本地 HLS 代理将加密视频分片误判为播放列表的问题：播放会话存在密钥或分片 URL 查询包含 `m3u8` 时，旧逻辑会把分片作为文本播放列表读取并返回 HTTP 502。现在按媒体类型、路径扩展名和内容识别播放列表，分片保留原字节、Range、Content-Type 与状态码。已用合成分片复现旧版错误并通过修复后的定向回归，覆盖加密分片、含播放列表关键词的查询参数、GET/HEAD 和无扩展名播放列表；实际黄豆样本经过本地代理后可由 FFprobe 识别 H.264 视频与 AAC 音频。播放器错误事件也会记录到“播放调试日记”。已完成 [0.2.66+73 全站源 iOS 云构建](https://github.com/GaodYang/guoapp/actions/runs/36706302759)，产物哈希与 Payload 结构校验通过；2026-09-30 的 Xcode 真机日志确认旧版 0.2.64+71 连续收到本机分片 HTTP 502，修复后真机效果待确认。

`0.2.67+74` 继续修正入口媒体判断：MP4 地址路径或查询参数含 `hls` / `m3u8` 不再直接认定为播放列表。新增合成 MP4 定向回归覆盖 GET/HEAD、Range、媒体字节和类型；无扩展名 HLS 仍按响应内容识别。封面点击现在立即打开“正在进入播放”页面，详情失败后可重试；分集目录请求超时显示中文提示，播放日记新增详情加载起止、耗时和失败原因。加载提示改善等待反馈，不能保证源站请求不再超时。 已完成 [0.2.67+74 全站源 iOS 云构建](https://github.com/GaodYang/guoapp/actions/runs/36709041853)（源码提交 `eaed65c`）；下载产物的 GitHub SHA256、包内 SHA256、`Payload/Runner.app`、Mach-O、版本 `0.2.67` 与构建号 `74` 已核对。用户已用原 Apple ID 自签覆盖安装，2026-09-30 通过 Xcode 设备查询确认手机版本为 `0.2.67`、构建号 `74`，并已连接新版控制台。真机确认点击封面会显示加载提示；黄果 AI 的 `117`、`2181`、`2833`、`12` 在手机直连时加载分集目录约 60 秒超时，同一批剧集在 Mac 生产核心复验均为 2–4 秒成功。用户随后开启手机代理并确认原先卡住的内容恢复快速加载，但随后反馈黄果 AI 实际错播其他剧，且本机媒体代理出现 Connection refused；加载恢复不能作为正确播放验收。 本版启动日记及关于页仍有写死的旧版本文字，设备系统版本查询结果以 `0.2.67+74` 为准，显示文字的统一读取在 0.2.68+75 修正。黄豆 188 集样本目录实际加载为 315 ms，不能把故障单独归因于集数。该结果只覆盖本次反馈样本，不能作为全部站源、全部分集的播放验收。

`0.2.68+75` 修复黄果 AI 错播：旧分集解析收集整个页面的 `/video/` 链接，将推荐区的其他剧误列为当前剧的分集，导致 `117`、`2181`、`2833` 均首先播放推荐视频 `6756`。现在只收录同一站点、同一剧集 ID 的分集，按 URL 的真实集数排序并校验 `data-ep-id`；取流前后校验页面剧集及分集身份，结构化播放数据的 ID 必须匹配，指定集数缺失时不再退回第一集或全页推荐媒体。旧任务中混入的其他剧页面会被拒绝，需要重新打开正确剧集。生产核心复验三部剧目录分别为 42、11、6 集，每部前两集的页面身份匹配，六个媒体地址互不相同且本机代理播放列表返回 HTTP 200；只核对文字与播放列表，未检查站源图片，真机画面仍待安装本版后确认。针对推荐区污染、真实集数排序、旧任务、错剧/错集、重定向、缺失集数和两部剧详情到取流的合成回归及相关竞态检查通过。

本版同时修复本机媒体代理关闭后继续返回旧端口的问题：每次创建播放会话前检查服务状态，发现失效后重建；服务启动和退出原因记入原生诊断日志。已用合成数据复现关闭服务后的 Connection refused，并验证连续三次重建可重新取得播放列表；真机服务退出的具体原因尚未确认，不能将恢复机制等同于所有断流问题已解决。启动日记和关于页统一显示 `0.2.68+75` 对应版本信息。已完成 [0.2.68+75 全站源 iOS 云构建](https://github.com/GaodYang/guoapp/actions/runs/36727374154)（源码提交 `a32bc82`），构建号常量已避开 Flutter services 同名符号；构建脚本增加 `--verbose` 详细诊断。GitHub 产物归档 SHA256 为 `8977f087a1aa8d0ce29fd28ba3e56492d1c86950072791068c11f75b1da069c8`，未签名 IPA SHA256 为 `2b5182b2453f59d12b704260fb1031d350eb17d19fbf29fd8d1438cc74f51073`；包内校验和、Payload 结构、Mach-O、版本 `0.2.68`、构建号 `75` 均通过核对。2026-09-30 用户补充闲置一段时间后复发 Connection refused，重启 App 可恢复，与本版媒体服务恢复的修复范围一致；用户已用原 Apple ID 自签覆盖安装；Xcode 设备查询确认手机版本为 `0.2.68`、构建号 `75`，新版原生诊断已记录本机媒体代理启动。具体真机退出原因、正确内容及闲置恢复效果仍待用户复验。

2026-09-30 用生产 Go 核心检测全部 11 个站源的文字 API、媒体开头和实际本机代理码流。黄豆、黄果视频、黄果旧版、青空、鬼片各两个样本通过 FFprobe 码流识别；黄果 AI 当时的两个样本虽然识别出码流，但未验证内容身份，后续已确认旧解析误收推荐视频，不能作为该站源正确播放的证据；剧果取流或探测超时，野果先通过连接检测、后续目录请求超时，韩小圈详情请求超时、复查一个分集播放地址返回 HTTP 404；帝果前两个样本没有可用于检测的免费分集。红果连接检测通过，初测两个样本本机代理误判播放列表；补充入口判断修复后，同样两个样本均返回 video/mp4 并识别出 HEVC 视频与 AAC 音频。抽样失败不能据此判定整个站源失效，连接检测通过也不等同于真机播放验收；未检查或下载站源图片。

该次云构建的原始 IPA 打包遗漏 `Payload/` 顶层目录。已修正 `build_ios.py` 的 `ditto --keepParent` 参数，并用构建产物重新打包；重新打包后确认 `Payload/Runner.app` 结构正确，157 个应用文件内容与云端产物一致。该脚本修复已纳入 `0.2.65+72` 源码；旧构建签名前应使用修复后的本地 IPA。

推送 `main` 且 android / ios / windows 全部构建成功时，自动创建 / 更新 GitHub Release（tag `app-v{version}`）。发布新版本前需先在 `pubspec.yaml` 提升 `version`，否则会覆盖同名 tag 的 Release。

Android 正式发布使用同一签名并递增构建号，在仓库 Secrets 配置：`ANDROID_KEYSTORE_BASE64`、`ANDROID_KEYSTORE_PASSWORD`、`ANDROID_KEY_ALIAS`、`ANDROID_KEY_PASSWORD`。未配置时生成 debug 签名预览包。

本地不入库的 `android/key.properties`：

~~~properties
storeFile=/absolute/path/zhenguojian-release.jks
storePassword=你的密码
keyAlias=zhenguojian
keyPassword=你的密码
~~~

## 开发与构建

Flutter `3.47.4`、Dart `3.12+`、Go `1.24.1+`、Python `3.10+`。Android 需要 JDK 17、SDK 36、NDK `28.2.13676358`；Windows 需要 Visual Studio C++ 桌面组件及 MinGW-w64 x64；iOS 需要 macOS、完整 Xcode 和 CocoaPods。

构建脚本对子进程默认设置 `GOPROXY=https://goproxy.cn,direct`、`GOSUMDB=off`，同名环境变量可覆盖。

~~~sh
python3 scripts/build_android.py                 # 红果鉴
python3 scripts/build_android.py --all-sources   # 真果鉴
python3 scripts/build_android.py --abi arm64-v8a
python3 scripts/build_android.py --cn-mirrors    # 国内镜像
~~~

~~~powershell
.\scripts\build_windows.ps1
.\scripts\build_windows.ps1 -AllSources
.\scripts\build_windows.ps1 -ChinaMirrors
~~~

~~~sh
python3 scripts/build_ios.py
python3 scripts/build_ios.py --all-sources
python3 scripts/build_ios.py --core-only [--simulator]
python3 scripts/build_ios.py --export-options /path/to/ExportOptions.plist
~~~

产物在 `dist/android`、`dist/windows`、`dist/ios`，红果版以 `hongguojian-` 开头，全站源版以 `zhenguojian-` 开头。

首次 Android 调试先编译对应架构核心，再运行：

~~~sh
python3 scripts/build_native.py --platform android --abi arm64-v8a
flutter pub get --enforce-lockfile
flutter run
~~~

调试全站源版：先给 `build_native.py` 加 `--all-sources`，再 `flutter run --dart-define=ALL_SOURCES=true`；iOS 对应 `build_ios.py --core-only --all-sources`。脚本会同步设置 Dart 常量和 Go 编译参数，应用启动时校验二者一致，避免混装原生库。

播放器使用 [media_kit](https://github.com/media-kit/media-kit) / libmpv，合并和导出使用 [FFmpegKit min-gpl](https://github.com/sk3llo/ffmpeg_kit_flutter)（含 GPL 媒体组件）。FFmpegKit 不参与正常播放或下载的转码。

### 集中检查与真机回归

~~~sh
python3 -m unittest discover -s scripts -p 'test_*.py'
dart format --output=none --set-exit-if-changed lib test integration_test test_driver
dart analyze --fatal-infos lib test integration_test test_driver
flutter test --dart-define=DISABLE_REMOTE_IMAGES=true
flutter test --dart-define=DISABLE_REMOTE_IMAGES=true --dart-define=ALL_SOURCES=true
cd native
go test -race ./...
go test -race -ldflags="-X duanjuapp/native/core.buildAllSources=true" ./...
~~~

Android 设备回归（连接并授权 USB 调试、保持解锁）：

~~~sh
python3 scripts/create_test_media.py
python3 scripts/serve_test_media.py
adb reverse tcp:38473 tcp:38473
flutter drive --driver=test_driver/playback.dart --target=integration_test/playback_test.dart \
  --dart-define=DISABLE_REMOTE_IMAGES=true --dart-define=FIXTURE_BASE_URL=http://127.0.0.1:38473
~~~

结果在 `build/device-test/results/`；结束后停服务并 `adb reverse --remove tcp:38473`。

### 源码同步

~~~sh
python3 scripts/finish_task.py --message "本次实际完成的变更"
python3 scripts/sync_source.py --check
~~~

脚本只同步纯源码到同级 `../guoapp`，并生成 `真果·鉴-YYYYMMDDHHMM.zip` 源码压缩包；不执行 Git 提交、分支或推送。

## 目录结构

| 目录 | 内容 |
| --- | --- |
| `lib` | 页面、播放器、本地用户、FFI、下载和媒体处理 |
| `native/core`、`native/bridge` | 独立站源核心、缓存、下载、目录迁移及 C ABI |
| `android`、`windows`、`ios` | 平台工程与必要资源 |
| `assets/video_enhancement`、`packages/media_kit_libs_windows_video` | 增强 Shader 与许可、固定 Windows 媒体依赖插件 |
| `scripts`、`.github/workflows` | 构建、签名、验证、同步和版本快照 |
| `test`、`integration_test` | 自动化与设备回归 |

## 站源开发约定

站源是 Go 原生 provider（`native/core/provider_*.go`），不是运行期加载的 Python 源。新增站源需接入：`provider_huangguo.go`（常量 / 白名单 / `canonicalProviderSource` / `GetHuangguoChapters`）、`provider_media.go`（baseURL / host 反查 / 播放分派）、`app_categories.go`、`app_cover_metadata.go`、`app_runtime.go`（Config 字段 + 目录 / 搜索 / 详情分派 + 空页放行），并在 `lib/models.dart` 登记 `SourceSite`。

注意：Go RE2 正则**不支持 lookahead**。解析多线路播放列表时不能用 `(?=...)` 做分段，否则非捕获组会消耗下一段开头；应改用显式字符串截断（从容器标记之后查找下一段标记）。
