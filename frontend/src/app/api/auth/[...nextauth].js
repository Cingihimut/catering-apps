// pages/api/auth/[...nextauth].js
import NextAuth from "next-auth";
import Providers from "next-auth/providers";

export default NextAuth({
  providers: [
    Providers.Credentials({
      credentials: {
        username: { label: "Username", type: "text" },
        password: { label: "Password", type: "password" },
      },
      authorize: async (credentials) => {
        const response = await fetch("http://localhost:8080/api/users/login", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(credentials),
        });

        if (response.ok) {
          const data = await response.json();
          return Promise.resolve(data);
        } else {
          return Promise.resolve(null);
        }
      },
    }),
  ],
  callbacks: {
    jwt: async (token, user) => {
      // Store the token in the user object
      if (user) {
        token.accessToken = user.token;
      }
      return token;
    },
    session: async (session, token) => {
      // Store the token in the session
      session.accessToken = token.accessToken;
      return session;
    },
  },
  pages: {
    signIn: "/auth/signin", // Customize the sign-in page URL
  },
  /* Other NextAuth.js options */
});
