import MainMenu from "../../components/MainMenu";
import { getClothingItems } from "@/services/getClothing";

export default async function MainMenuPage() {
  const clothingItems =
    await getClothingItems();

  return (
    <MainMenu
      clothingItems={clothingItems}
    />
  );
}