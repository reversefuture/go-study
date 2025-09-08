JWT（JSON Web Token）本身是一种开放标准（RFC 7519），用于在各方之间安全地传输信息作为 JSON 对象。它通过使用**加密算法**来确保信息的完整性和/或机密性。JWT 的实现主要依赖于以下几类算法，具体取决于使用场景（是否需要签名或加密）：

---

### 一、签名算法（最常见，用于保证完整性）
JWT 通常使用 **数字签名** 来验证发送方身份并确保数据未被篡改。常见的签名算法属于 **JWS（JSON Web Signature）** 标准。

#### 1. HMAC 算法（对称加密）
- **HS256**：HMAC + SHA-256（最常用）
- HS384：HMAC + SHA-384
- HS512：HMAC + SHA-512

> 特点：使用同一个密钥进行签名和验证，适合服务端内部使用。

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

#### 2. RSA 算法（非对称加密）
- **RS256**：RSA + SHA-256（非常常见，适合分布式系统）
- RS384：RSA + SHA-384
- RS512：RSA + SHA-512

> 特点：使用私钥签名，公钥验证，适合开放系统（如 OAuth 2.0、OpenID Connect）。

#### 3. ECDSA 算法（椭圆曲线数字签名算法）
- **ES256**：ECDSA with P-256 + SHA-256
- ES384：ECDSA with P-384 + SHA-384
- ES512：ECDSA with P-521 + SHA-512

> 特点：安全性高，密钥短，适合移动端或资源受限环境。

---

### 二、加密算法（用于保证机密性）
如果需要对 JWT 的内容加密（即不让中间方看到 payload），可以使用 **JWE（JSON Web Encryption）**。

常用加密算法包括：
- **AES**（Advanced Encryption Standard）：如 A128GCM、A256CBC-HS512
- **RSA 加密密钥**：如 RSA-OAEP、RSA1_5
- 密钥交换机制：如 ECDH-ES

> 注意：JWE 使用较少，大多数 JWT 场景只做签名，不加密（payload 是 Base64 编码可读的）。

---

### 三、JWT 的结构
JWT 由三部分组成，用 `.` 分隔：
```
Header.Payload.Signature
```

- **Header**：指定算法（`alg`）和类型（`typ`）
- **Payload**：包含声明（claims），如用户 ID、过期时间等
- **Signature**：使用指定算法对前两部分签名生成

---

### 四、常见算法总结表

| 算法 | 类型 | 是否对称 | 常见用途 |
|------|------|----------|----------|
| HS256 | HMAC | 是 | 内部服务认证 |
| RS256 | RSA | 否 | OAuth、OpenID |
| ES256 | ECDSA | 否 | 移动端、高安全场景 |
| PS256 | RSA-PSS | 否 | 更安全的 RSA 变种 |
| none | 无 | — | 不推荐，不安全 |

> ⚠️ `alg: none` 表示无签名，存在严重安全风险，不推荐使用。

---

### 五、选择建议
- 内部微服务：使用 **HS256**
- 开放平台、第三方登录：使用 **RS256**
- 高安全或移动端：考虑 **ES256**

---

### 示例（HS256 签名流程）：
1. 编码 Header 和 Payload 为 Base64Url
2. 拼接：`base64(header) + "." + base64(payload)`
3. 使用密钥对拼接结果进行 HMAC-SHA256 签名
4. 将签名结果 Base64Url 编码，作为第三段

---

✅ 总结：  
**JWT 本身不是一种算法，而是一种格式标准，其安全性依赖于底层的签名或加密算法**。最常用的算法是 **HS256** 和 **RS256**。

如需更高安全或特定场景，可选用 ES256 或 JWE 加密方案。