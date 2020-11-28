# go-encoder-playground

This project was created to explore idiosyncrasies of the inbuilt [GoLang Encoding Package](https://golang.org/pkg/encoding/). Google project zero recently highlighted an interesting behaviour in their blog post [Enter the Vault: Authentication Issues in HashiCorp Vault](https://googleprojectzero.blogspot.com/2020/10/enter-the-vault-auth-issues-hashicorp-vault.html). Highlighting the decoder silently ignores non XML content before and after the expected XML root. For reference, the following paragraph from the blog post describes the problem well:

> For Vault, this turns into a security issue due to a somewhat surprising feature of the Go XML decoder: The decoder silently ignores non XML content before and after the expected XML root. This means that calling parseGetCallerIdentityResponse with a (JSON encoded) server response such as ‘{“abc” : “xzy<GetCallerIdentityResponse></GetCallerIdentityResponse>}’ will succeed and return an (empty) CallerIdentityResponse structure.

I’m interested in exploring this for myself and gaining hands on experience.