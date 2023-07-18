FROM node:lts
WORKDIR /app
COPY package.json .
RUN npm i
COPY ./FrontHotel .
## EXPOSE [Port you mentioned in the vite.config file]
EXPOSE 8000
CMD ["npm", "run", "dev"]