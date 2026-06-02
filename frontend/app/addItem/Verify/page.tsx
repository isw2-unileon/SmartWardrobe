import AddItemVerify from "@/components/AddItemVerify";

export default async function VerifyPage({
  searchParams,
}: {
  searchParams: Promise<{ imageUrl?: string }>;
}) {
  const params = await searchParams;

  return <AddItemVerify imageUrl={params.imageUrl || ""} />;
}
