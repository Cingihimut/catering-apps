const jwt = require("jsonwebtoken");

const SECRET = process.env.JWT_SECRET;
const decodeToken = (token) => {
  try {
<<<<<<< HEAD
    console.log(SECRET);
    const decoded = jwt.verify(token, SECRET);
=======
    console.log(token);
    const decoded = jwt.verify(token, process.env.JWT_SECRET);
>>>>>>> main
    return decoded;
  } catch (error) {
    console.error("Gagal mendekode token:", error.message);
    return null;
  }
};

export default decodeToken;
