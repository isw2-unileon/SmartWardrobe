import ModifyItemForm from "@/components/ModifyItemForm";

export default async function ModifyItemPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;

  return (
    <ModifyItemForm
      itemId={Number(id)}
    />
  );
}