"use server";

import { createClient } from "@/utils/supabase/server";
import { cookies } from "next/headers";

export async function uploadImage(
  formData: FormData
) {
  const file = formData.get(
    "file"
  ) as File;

  const cookieStore = await cookies();
  const supabase = createClient(cookieStore);

  const fileName = `${Date.now()}-${file.name}`;

  const { data, error } =
    await supabase.storage
      .from("wardrobe-images")
      .upload(fileName, file);

  if (error) {
    throw new Error(error.message);
  }

  const {
    data: { publicUrl },
  } = supabase.storage
    .from("wardrobe-images")
    .getPublicUrl(fileName);

  return publicUrl;
}