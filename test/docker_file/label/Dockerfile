FROM alpine:latest

LABEL author="tirtahakimpambudhi"
LABEL github="https://github.com/tirtahakimpambudhi" dockerhub="https://hub.docker.com/u/tirtahakimpambudhi" email="tirtanewwhakim22@gmail.com"

RUN echo "Start"
RUN mkdir -p app/log
RUN echo "Saya Tirta Hakim Pambudhi" > app/log/biodata.log && \ 
    echo "Email Saya tirtanewwhakim22@gmail.com" >> app/log/biodata.log && \
    echo "Github : https://github.com/tirtahakimpambudhi" >> app/log/biodata.log && \
    echo "Dockerhub : https://hub.docker.com/u/tirtahakimpambudhi" >> app/log/biodata.log 
RUN echo "Finish"
CMD cat app/log/biodata.log