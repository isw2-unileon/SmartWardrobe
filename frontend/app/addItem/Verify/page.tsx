import AddItemVerify from "@/components/AddItemVerify";

export default async function VerifyPage({
  searchParams,
}: {
  searchParams: Promise<{
    imageUrl?: string;
    color?: string;
    style?: string;
    type?: string;
  }>;
}) {

  const params =
    await searchParams;

  return (
    <AddItemVerify
      imageUrl={
        params.imageUrl || ""
      }
      predictedColor={
        params.color || ""
      }
      predictedStyle={
        params.style || ""
      }
      predictedType={
        params.type || ""
      }
    />
  );
}