"use client";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { Provider } from "jotai";

const qc = new QueryClient();

export default function Providers({ children }: { children: React.ReactNode }) {
  return (
    <Provider>
      <QueryClientProvider client={qc}>
        {children}
        <ReactQueryDevtools initialIsOpen={false} client={qc} />
      </QueryClientProvider>
    </Provider>
  );
}
