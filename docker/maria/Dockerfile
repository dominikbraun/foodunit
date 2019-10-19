#
# FoodUnit 3 MariaDB image
#

FROM mariadb:latest

# PACKAGES defines the apt packages to get installed.
ENV PACKAGES openssh-server openssh-client

RUN apt-get update && apt-get install -y ${PACKAGES}
RUN sed -i 's|^#PermitRootLogin.*|PermitRootLogin yes|g' /etc/ssh/sshd_config; \
    echo "root:root" | chpasswd