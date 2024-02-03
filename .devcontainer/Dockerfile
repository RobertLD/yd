FROM mcr.microsoft.com/oss/go/microsoft/golang:1.21

### Section to create a non-root user with sudo priveleges ########
# USER, USER_UID, USER_GID is specified in devcontainer.json
ARG USERNAME=vscode
ARG USER_UID=1001
ARG USER_GID=${USER_UID}


# # Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# # ********************************************************
# # * Anything else you want to do like clean up goes here *
# # ********************************************************

# Install the Microsoft ODBC driver for SQL Server
# https://docs.microsoft.com/en-us/sql/connect/odbc/linux-mac/installing-the-microsoft-odbc-driver-for-sql-server?view=sql-server-ver15
ARG DEBIAN_FRONTEND=noninteractive

RUN curl https://packages.microsoft.com/keys/microsoft.asc | apt-key add -
# Debian 10
RUN curl https://packages.microsoft.com/config/debian/11/prod.list > /etc/apt/sources.list.d/mssql-release.list

RUN apt-get update && \
    ACCEPT_EULA=Y apt-get install -y msodbcsql17 && \
    ACCEPT_EULA=Y apt-get install -y mssql-tools && \
    . ~/.bashrc && \
    apt-get -y install \
    unixodbc-dev \
    openjdk-17-jre \
    krb5-user \
    sudo \
    fonts-powerline \
    wget

ENV HOME /home/${USERNAME}
ENV SHELL /bin/bash
