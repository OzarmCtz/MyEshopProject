FROM node:18-alpine

WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY . ./

# Build pour la prod (pas de recompilation)
RUN npm run build

EXPOSE 3000

# Start en mode production
CMD ["npm", "start"]