import ModifyItemForm from "@/components/ModifyItemForm";
import { getClothingById } from "@/services/getClothingById";

export default async function ModifyItemPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;

  const item = await getClothingById(Number(id));

  return <ModifyItemForm item={item} />;
}
