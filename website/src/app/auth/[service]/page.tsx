type PageProps = {
  params: {
    service: string;
  };
};

export default function Page({ params }: PageProps) {
  return <div>{params.service}</div>;
}
