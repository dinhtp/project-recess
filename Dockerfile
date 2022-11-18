FROM ubuntu:22.04

EXPOSE 9090

ADD ./bin/project-recess /usr/bin/project-recess

CMD ["project-recess", "serve"]