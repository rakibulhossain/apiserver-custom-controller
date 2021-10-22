
vendor/k8s.io/code-generator/generate-groups.sh all \
github.com/Rakibul-Hossain/apiserver-custom-controller/pkg/generated \
github.com/Rakibul-Hossain/apiserver-custom-controller/pkg/apis \
customcontroller:v1 \
--output-base "${GOPATH}/src" \
--go-header-file "hack/boilerplate.go.txt"