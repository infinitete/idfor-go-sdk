# CTID 操作

### 1. 绑定关联
```
sdk.Native.OntId.AssignCtidPID(identity *sdk.Identity, acc sdk.Account, bas, pid, pass []byte) (common.Uint256, error)
```
注意：一个数字身份可以关联多个bas，但是每个bas只能关联一个pid，pid会使用账户的公钥加密后上链。上链成功后链上会有事件提示，格式如下：
```json
{
  "TxHash": "4c7f8171142ebf56783adcd55a31020ad09b0f00dfea9f4a546c49f777a993d1",
  "State": 1,
  "GasConsumed": 0,
  "Notify": [
    {
      "ContractAddress": "0300000000000000000000000000000000000000",
      "States": [
        "AssumeCtidPid",
        "did:idfor:TKc4VfPTBGZGJp5ybsVfCnELKCybsYQzSR",
        "fecredBas-475",
        "ed0a5f0e59675786965364619f17320e838b8f6fbdb8f9a2c866f7f5870e8c2a67fc22797d2f8139efcfadef1e33e873ced5993576fe563db98a43e3c97d6650a9e2eaa867649ab5d04a946002756d5f9598ef348f3031cba7a886f1df7dda476916d40598ddd75b5432b37cdcc0",
        "09008af061444d622899c9b01b91a85c366dc6dacdff041a56c1527446f956bfc313c35e4dbd28fd0d352cf5d0bbb0a45a81fd88267b5c6e1a4a0d89055b3723ba37"
      ]
    }
  ]
}
```
`State`数组中的第一个是操作函数，第二个元素是操作的数字身份，第三个是bas标识符，第四个是加密后的pid，第五个是该数字身份私钥对加密后pid的签名。

### 2. 获取自己的PID
```
sdk.Native.OntId.GetCtidPID(identity *Identity, acc *Account, bas, pass []byte) (string, error)
```
如果正确，返回的地一个参数是pid，否则返回错误信息

### 3. 获取他人的EPID
```
sdk.Native.OntID.GetCtidEPID(id, bas []byte) (string, string, string, error)
```
参数`id`是需要获取的数字身份，`bas`是bas标识符，如果正确，第一个返回值是`bas`标识符，第二个返回值是加密的`pid`(`epid`)，第三个是签名`signature`，第四个为空，否则第四个为错误信息。

### 4. 撤销PID
```
sdk.Native.OntId.RevokeCtidPID(identity *Identity, acc *Account, bas, pass []byte) (common.Uint256, error)
```
