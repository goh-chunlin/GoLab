FROM scratch

EXPOSE 80

COPY GoLab /
COPY public public
COPY templates templates

ENV APPINSIGHTS_INSTRUMENTATIONKEY '' \
    CONNECTION_STRING '' \
    OAUTH_CLIENT_ID '' \
    OAUTH_CLIENT_SECRET '' \
    COOKIE_STORE_SECRET '' \
    OAUTH2_CALLBACK ''

CMD [ "/GoLab" ]