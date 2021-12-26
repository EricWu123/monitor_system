系统负载监控系统
Agent端：
（1）收集系统指标
①主机信息：主机名、操作系统信息
②网络信息：IP地址列表、网络读写字节/包个数
③CPU信息：CPU型号、CPU逻辑核数、CPU物理核心、CPU使用率、
④内存使用情况：总内存大小、实际内存大小、交换区大小
⑤硬盘使用情况：总硬盘大小、已使用硬盘大小
⑥参考库： github.com/shirou/gopsutil
（1）定时上报
①接口鉴权：将固定AppCode放在Header中，请求Header中添加的Authorization字段；配置Authorization字段的值为“APPCODE ＋ 半角空格 ＋APPCODE值”。例如 Authorization:APPCODE AppCode值
②每分钟上报一次系统指标结构化数据
Server端
（1）支持Agent上报数据接口
①接口鉴权：校验http header中的AppCode值是否一致
②上报数据持久化存储
（2）支持多用户登录
①用户信息需要做持久化，用户密码不能明文存放
②用户账户可以提前创建好，支持admin和guest
③除登录接口外，其它接口需要鉴权
（3）提供查询上报数据接口
①查询维度：主机名；操作系统类型 Linux、Windows、Mac；上报时间段
（4）拓展（选做）
①记录用户查询操作：用户ID、操作时间、查询参数
②提供审计接口：仅支持admin查看、查看所有用户的操作记录

关注代码质量：
①API设计符合规范
②模块有良好的封装和抽象