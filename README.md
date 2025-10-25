
### 来自openRTB

数据格式: 压缩 JSON、ProtoBuf、Avro 等，这些形式在传输时间和带宽方面可能更高效


数据编码: 压缩交易所和竞价者之间传输的数据可以带来很大好处。压缩可以显著减小传输数据的大小，从而节省交易所和竞价者的网络带宽。为了充分实现这些节省，交易所发送的出价请求和竞价者返回的出价响应都应该启用压缩功能
客户端使用: Accept-Encoding: gzip  (表示我支持gzip 压缩格式)
服务器使用: Content-Encoding: gzip (表示看到客户端支持我服务器响应一个压缩格式的内容)

openRTB版本HTTP头部
x-openrtb-version:2.6

