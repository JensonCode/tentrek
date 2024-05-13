import { useState, useEffect } from "react";

export const useCounter = (second: number) => {
  const [seconds, setSeconds] = useState(second);

  useEffect(() => {
    if (seconds === 0) return;

    const timer = setInterval(() => {
      setSeconds((prev) => prev - 1);
    }, 1000);

    return () => clearInterval(timer);
  }, [seconds]);

  return { seconds };
};
