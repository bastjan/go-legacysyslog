# go-legacysyslog

![](https://github.com/bastjan/go-legacysyslog/workflows/Go/badge.svg)

⚠️ Work in progress ⚠️

A parser for BSD (rfc3164) syslog. Aims to be compatible with the syslog-ng legacy parser.

## Known Incompatibilities to Syslog-NG

- [ ] Handling of misformed priority. `<<` results in a parsing error in Syslog-NG. We accept it as **tag**. `<` is seen as kernel emergency in Syslog-NG, we see it as tag.
- [ ] Handling of timestamps. Syslog-NG allows space instead of zeros. `Dec  1 09: 5: 2` is valid.
- [ ] Check hostname and validate UTF8 options are not implemented.
- [ ] Ignoring `last message repeated` / `forwarded from` messages is not implemented.

## Testing

```sh
make test
```
