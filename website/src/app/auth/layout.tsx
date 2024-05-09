import Logo from "@/components/logo";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="h-screen">
      <div className="flex h-[10dvh] items-center justify-center">
        <Logo />
      </div>

      <main className="flex h-[80dvh] items-center justify-center px-4 py-12 sm:px-6">
        {children}
      </main>
    </div>
  );
}
