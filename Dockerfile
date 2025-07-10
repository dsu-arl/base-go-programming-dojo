FROM golang:latest

USER root

# RUN <<EOF
# 	if id ubuntu; then userdel -f -r ubuntu; fi
# 	useradd -s /bin/bash -m hacker
# 	passwd -d hacker

# 	ln -s /etc/bash.bashrc /etc/bashrc
# EOF


RUN <<EOF
	apt update && apt install git wget build-essential zlib1g-dev libssl-dev libncurses5-dev libnss3-dev libreadline-dev libffi-dev libsqlite3-dev libbz2-dev -y
    cd /usr/src
    wget https://www.python.org/ftp/python/3.12.0/Python-3.12.0.tgz
    tar -xzf Python-3.12.0.tgz
EOF

RUN <<EOF
    cd /usr/src/Python-3.12.0
    ./configure --enable-optimizations
    make -j$(nproc)
    make altinstall
    update-alternatives --install /usr/bin/python3 python3 /usr/local/bin/python3.12 1
EOF

RUN <<EOF
    rm -rf /usr/src/Python-3.12.0
    rm /usr/src/Python-3.12.0.tgz
    apt remove wget build-essential zlib1g-dev libssl-dev libncurses5-dev libnss3-dev libreadline-dev libffi-dev libsqlite3-dev libbz2-dev -y
EOF

ADD --chown=0:0 --chmod=6755 http://github.com/pwncollege/exec-suid/releases/latest/download/exec-suid /usr/bin/exec-suid
# RUN chmod 6755 /usr/bin/exec-suid