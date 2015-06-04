FROM busybox:ubuntu-14.04

MAINTAINER klumhru@gmail.com (Högni Gylfason)

ADD intego /usr/sbin/intego

EXPOSE 8080

CMD ["/usr/sbin/intego"]