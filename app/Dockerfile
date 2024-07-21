FROM node:21-alpine3.19 AS base
WORKDIR /usr/src/app

ARG VITE_API_URL

FROM base AS install
RUN mkdir -p /temp/dev
COPY package.json yarn.lock /temp/dev/
RUN cd /temp/dev && yarn install --frozen-lockfile

FROM base AS release
COPY --from=install /temp/dev/node_modules node_modules
COPY --from=install /temp/dev/package.json .
COPY . .

EXPOSE 5173/tcp
ENTRYPOINT [ "yarn", "dev", "--host" ]
# Ideally for a real production build you'd build the static pages and serve it using something like Nginx