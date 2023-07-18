FROM node:lts
WORKDIR /app
COPY package.json .
RUN npm i
COPY ./FrontHotel .
EXPOSE 8000
CMD ["npm", "run", "dev"]