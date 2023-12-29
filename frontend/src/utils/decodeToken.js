const jwt = require("jsonwebtoken");

const SECRET = process.env.JWT_SECRET || "CATERING";
const decodeToken = (token) => {
  try {
    console.log(SECRET);
    const decoded = jwt.verify(token, SECRET);
    return decoded;
  } catch (error) {
    console.error("Gagal mendekode token:", error.message);
    return null;
  }
};

export default decodeToken;
