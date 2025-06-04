import AsyncStorage from "@react-native-async-storage/async-storage";
import axios, { AxiosError, AxiosInstance, AxiosResponse } from "axios";

// const url = Platform.OS === "android" ? "http://10.0.2.2:8081" : "http://127.0.0.1:3000"
// const url = "https://9c05-2409-40c4-309-e6d4-ac69-1e3-8775-f01f.ngrok-free.app" //=> by ngrok api gateway
// Paste your own api key , I have not deployed the backend yet
const Api: AxiosInstance = axios.create({ baseURL: url + "/api" });

Api.interceptors.request.use(async config => {
  const token = await AsyncStorage.getItem("token");

  if (token) config.headers.set("Authorization", `Bearer ${token}`);

  return config;
});

Api.interceptors.response.use(
  async (res: AxiosResponse) => res.data,
  async (err: AxiosError) => Promise.reject(err)
);

export { Api };

