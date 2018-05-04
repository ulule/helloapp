FROM debian:stretch-slim

ADD bin/helloapp /helloapp

CMD ["/helloapp"]
