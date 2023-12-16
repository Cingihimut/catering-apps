const jwt = require("jsonwebtoken");

const decodeToken = (token) => {
  try {
    console.log(token);
    const decoded = jwt.verify(token, "SECRET");
    return decoded;
  } catch (error) {
    console.error("Gagal mendekode token:", error.message);
    return null;
  }
};

export default decodeToken;