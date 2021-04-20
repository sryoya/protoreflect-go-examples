# protoreflect-go-examples

- [日本語版はこちら(Japanese Description is here)](#protoreflect-go-examples (日本語版, Description by Japanese))

## About


## Appendix

### 1. String with custom regulations

[proto defenition](To be fiiled)
[implementation](To be fiiled)

It has custom rules for String fields like a length or a format.
The way to use is simple. Set the options to String fields in `.proto` and use the [`Validation`](To Be Filled) func to the concreate Go struct mapped with Proto Message.

#### Example

```.proto

```

```.go

```


### 2. Random Message generation for Protol Buffers Message
[implementation](To be fiiled)
It embeds values to [proto.Message]() with its values filled randomly.


#### Example

```.proto

```

```.go

Also, it can return [`dynamicpb`] message from [protoreflect.MessageDescriptor]
So, you can use it without having the concreat Go type messages.

```.go

```


# protoreflect-go-examples (日本語版, Description by Japanese)

## About
Go Protocol Buffers API v2のリフレクション機能の使用例になります。
[Go Conference Tokyo](https://gocon.jp/) 2021 Springで公開されたものなります。
[スライド](To be filled)はこちら

## Background

Protocol Buffers API v2が2020年にリリースされました。v2では、コード生成のライブラリなど開発者のインターフェースに関わる変更に加えて、Reflectionのサポートが提供されるようになりました。このReflection機能は、GoのビルトインのReflectionのように、Protocol Buffersのメッセージを動的に操作・参照することを可能にしてくれます。
ここでは、その使用例を２つほど取り上げています。

## Appendix

### 1. String with custom regulations
proto defenition
implementation

It has custom rules for String fields like a length or a format.
How to use is so simple. Set the options to String fields in `.proto` and call `Validation` to the concreate Go struct mapped with Proto Message.

#### Example

```.proto

```

```.go

```


### 2. Random Message generation for Protol Buffers Message
implementation

#### Example

```.proto

```

```.go

```
