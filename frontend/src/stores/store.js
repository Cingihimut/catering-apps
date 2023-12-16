import { create } from "zustand";

const useStore = create((set) => ({
  user: {
    email: null,
    id: null,
  },
  setUser: ({ email, id }) => set({ user: { email, id } }),
}));

export default useStore;
