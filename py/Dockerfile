FROM python:3-slim AS build-env
ADD . /app
WORKDIR /app
RUN pip install -r requirements.txt

FROM gcr.io/distroless/python3
ENV VIRTUAL_ENV=/opt/venv
RUN python3 -m venv $VIRTUAL_ENV
ENV PATH="$VIRTUAL_ENV/bin:$PATH"

COPY --from=build-env /app /app
WORKDIR /app
CMD ["catalog.py"]
