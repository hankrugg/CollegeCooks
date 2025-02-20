# CollegeCooks

## Overview  
This project is a web application built with Go and React that allows users to add the food in their pantry and receive recipe recommendations based on available ingredients. The goal is to help users reduce waste and make meal planning more efficient.

## Technologies Used  
- **Frontend:** React  
- **Backend:** Go  
- **Database:** PostgreSQL 

## Features  
- Add and manage pantry items  
- Generate recipe suggestions based on available ingredients  
- Search and filter recipes  

## Setup  
1. Clone the repository:  
   ```sh
   git clone https://github.com/hankrugg/CollegeCooks.git
   cd CollegeCooks
   ```
2. Install dependencies:  
   ```sh
   cd frontend  
   npm install  
   cd ../backend  
   go mod tidy  
   ```
3. Run the project:  
   ```sh
   cd backend  
   go run main.go  
   cd ../frontend  
   npm start  
   ```

## Future Enhancements  
- AI-powered recipe recommendations  
- Community recipe sharing  
- Mobile-friendly UI  

## License  
This project is open-source under the MIT License.  

---
*Contributions and feedback are welcome!*  
Following tutorial from https://dev.to/divrhino/build-a-rest-api-from-scratch-with-go-and-docker-3o54
