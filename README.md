# Generator

代码生成器

## Usage

在目录增加模版文件[mybatis-class.tmpl](./mybatis-class.tmpl)

增加配置文件`generator.json`

```json
{
  "dest": "./model",
  "pkg": "com.ljun20160606.mysql",
  "connectConfig": {
    "username": "username",
    "password": "password",
    "host": "host",
    "port": 3317,
    "scheme": "scheme"
  },
  "modelConfig": {
    "filename": "./mybatis-class.tmpl",
    "fileExtension": "java",
    "tableConfig": {
      "includes": ["table_name"],
      "excludes": []
    }
  }
}
```

```sh
# 运行命令
$ generator
```

在`./model`目录下会生对应实例下的所有表，其中单表结果如下

```java
package com.ljun20160606.mysql;

import java.io.Serializable;
import java.util.Date;
import lombok.Data;
import lombok.experimental.Accessors;

/**
 * table comment
 */
@Data
@Accessors(chain = true)
public class Account implements Serializable {

    private static final long serialVersionUID = 1L;

    /**
     * 自增主键 关联id
     */
    private Long id;

    ...
```