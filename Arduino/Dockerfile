FROM debian:bullseye

ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies
RUN apt-get update && apt-get install -y \
  build-essential \
  git \
  python3 \
  python3-pip \
  wget \
  unzip \
  curl \
  gcc-avr \
  avr-libc \
  avrdude \
  pkg-config \
  clang \
  cmake \
  && rm -rf /var/lib/apt/lists/*

# Install Arduino CLI
RUN curl -fsSL https://raw.githubusercontent.com/arduino/arduino-cli/master/install.sh | sh \
    && mv bin/arduino-cli /usr/local/bin/arduino-cli

# Build Google Test from source
RUN git clone https://github.com/google/googletest.git /usr/src/gtest && \
    cd /usr/src/gtest && \
    cmake . && \
    make && \
    cp lib/*.a /usr/lib

WORKDIR /workspace
COPY . /workspace

CMD ["bash"]