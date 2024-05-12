import axios, { AxiosResponse } from "axios";

type ServerError = {
  error: string;
};

export const getServerError = (err: unknown): Error => {
  if (axios.isAxiosError(err)) {
    const res: AxiosResponse<ServerError, unknown> | undefined = err.response;
    if (res) {
      return Error(res.data.error);
    }
  }

  return Error("Internal Server error");
};
