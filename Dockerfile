FROM eclipse-mosquitto:2.0.15 as cli

RUN mosquitto_ctrl dynsec init /mosquitto/config/dynamic-security.json admin 000000
# RUN mosquitto_ctrl -u admin -P 000000 dynsec createClient admin01 000000
COPY mosquitto.conf /mosquitto/config/mosquitto.conf