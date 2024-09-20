FROM ubuntu:latest

RUN apt-get update && \
  apt-get install -y \
  apt-transport-https \
  ca-certificates \
  curl \
  fonts-powerline \
  gcc \
  git-core \
  gnupg \
  golang-go \
  jq \
  make \
  pre-commit \
  python3 \
  python3-dev \
  python3-venv \
  python3-poetry \
  python3-pip \
  software-properties-common \
  wget \
  unzip \
  vim \
  zsh


# OPENTOFU
RUN chmod 755 /etc/apt/keyrings
RUN curl -fsSL https://get.opentofu.org/opentofu.gpg | tee /etc/apt/keyrings/opentofu.gpg >/dev/null
RUN curl -fsSL https://packages.opentofu.org/opentofu/tofu/gpgkey | gpg --no-tty --batch --dearmor -o /etc/apt/keyrings/opentofu-repo.gpg >/dev/null
RUN chmod a+r /etc/apt/keyrings/opentofu.gpg /etc/apt/keyrings/opentofu-repo.gpg

RUN echo   "deb [signed-by=/etc/apt/keyrings/opentofu.gpg,/etc/apt/keyrings/opentofu-repo.gpg] https://packages.opentofu.org/opentofu/tofu/any/ any main deb-src [signed-by=/etc/apt/keyrings/opentofu.gpg,/etc/apt/keyrings/opentofu-repo.gpg] https://packages.opentofu.org/opentofu/tofu/any/ any main" | tee /etc/apt/sources.list.d/opentofu.list > /dev/null
RUN chmod a+r /etc/apt/sources.list.d/opentofu.list

RUN apt-get update
RUN apt-get install -y tofu

# See https://github.com/tofuutils/pre-commit-opentofu?tab=readme-ov-file#terraftofu_fmtorm_fmt
# TOFU Lint
# RUN curl -L "$(curl -s https://api.github.com/repos/terraform-docs/terraform-docs/releases/latest | grep -o -E -m 1 "https://.+?-linux-amd64.tar.gz")" > terraform-docs.tgz && tar -xzf terraform-docs.tgz terraform-docs && rm terraform-docs.tgz && chmod +x terraform-docs && mv terraform-docs /usr/bin/
# RUN curl -L "$(curl -s https://api.github.com/repos/tenable/terrascan/releases/latest | grep -o -E -m 1 "https://.+?_Linux_x86_64.tar.gz")" > terrascan.tar.gz && tar -xzf terrascan.tar.gz terrascan && rm terrascan.tar.gz && mv terrascan /usr/bin/ && terrascan init
# RUN curl -L "$(curl -s https://api.github.com/repos/terraform-linters/tflint/releases/latest | grep -o -E -m 1 "https://.+?_linux_amd64.zip")" > tflint.zip && unzip tflint.zip && rm tflint.zip && mv tflint /usr/bin/
# RUN curl -L "$(curl -s https://api.github.com/repos/aquasecurity/tfsec/releases/latest | grep -o -E -m 1 "https://.+?tfsec-linux-amd64")" > tfsec && chmod +x tfsec && mv tfsec /usr/bin/
# RUN curl -L "$(curl -s https://api.github.com/repos/aquasecurity/trivy/releases/latest | grep -o -E -i -m 1 "https://.+?/trivy_.+?_Linux-64bit.tar.gz")" > trivy.tar.gz && tar -xzf trivy.tar.gz trivy && rm trivy.tar.gz && mv trivy /usr/bin
# RUN curl -L "$(curl -s https://api.github.com/repos/infracost/infracost/releases/latest | grep -o -E -m 1 "https://.+?-linux-amd64.tar.gz")" > infracost.tgz && tar -xzf infracost.tgz && rm infracost.tgz && mv infracost-linux-amd64 /usr/bin/infracost && infracost register
# RUN curl -L "$(curl -s https://api.github.com/repos/minamijoyo/tfupdate/releases/latest | grep -o -E -m 1 "https://.+?_linux_amd64.tar.gz")" > tfupdate.tar.gz && tar -xzf tfupdate.tar.gz tfupdate && rm tfupdate.tar.gz && mv tfupdate /usr/bin/
# RUN curl -L "$(curl -s https://api.github.com/repos/minamijoyo/hcledit/releases/latest | grep -o -E -m 1 "https://.+?_linux_amd64.tar.gz")" > hcledit.tar.gz && tar -xzf hcledit.tar.gz hcledit && rm hcledit.tar.gz && mv hcledit /usr/bin/


# BUF
ENV PREFIX="/usr/local"
## TODO: VERSION/VAR at top of file?
ENV VERSION="1.42.0"
RUN curl -sSL "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m).tar.gz" | tar -xvzf - -C "${PREFIX}" --strip-components 1


# Proto specific deps
# TODO ... Latest :-/
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest


## Set GOBIN on PATH
ENV GO_BIN "/root/go/bin"
# Add the Go bin directory to the PATH
ENV PATH="$PATH:$GO_BIN"


# GCLOUD-CLI
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | gpg --dearmor -o /usr/share/keyrings/cloud.google.gpg && apt-get update -y && apt-get install google-cloud-cli -y





ENV SHELL /bin/zsh


# POETRY STUFF
COPY pyproject.toml poetry.lock ./

## TODO: Consider a poetry.toml
RUN poetry config virtualenvs.create true

RUN poetry lock --no-update
RUN poetry install --no-interaction --no-ansi


# Store var
RUN ln -s $(poetry env info --path) /var/my-venv

# Styling
RUN wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -O - | zsh || true

## powerlevel10k
RUN git clone https://github.com/romkatv/powerlevel10k.git ~/.oh-my-zsh/custom/themes/powerlevel10k

RUN wget https://raw.githubusercontent.com/romkatv/dotfiles-public/master/.purepower -P $HOME

COPY .zshrc $HOME

# PATH/ENV
RUN echo 'source /var/my-venv/bin/activate' >> $HOME/.zshrc
