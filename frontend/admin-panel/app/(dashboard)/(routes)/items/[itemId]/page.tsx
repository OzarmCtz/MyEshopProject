import { Item, SubCategory } from '@/app/schema/schema';
import { ItemForm } from './components/item-form';
import axios from 'axios';
import { cookies } from 'next/headers';

const ItemPage = async ({ params }: { params: { itemId: number } }) => {
    let itemData: Item | null = null;
    let subCategoryData: SubCategory[] = [];
    let initialSubCategoryData: SubCategory | null = null;



    const cookieStore = cookies();
    const token = cookieStore.get('token')?.value;
    if (params.itemId) {
        try {
            const response = await axios.get(`${process.env.NEXT_PUBLIC_PV_PATH_API}/items/${params.itemId}?token=${token}`);
            if (response.status === 200) {
                itemData = response.data;
            }
        } catch (error) {
            console.error("Failed to fetch item:", error);
        }

        try {

            const subCategoryLinked = await axios.get(`${process.env.NEXT_PUBLIC_PV_PATH_API}/items/sub/category/link/${params.itemId}?token=${token}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            if (subCategoryLinked.status === 200) {
                initialSubCategoryData = subCategoryLinked.data;
            }
        } catch (error) {
        }


    }

    try {
        const subCategory = await axios.get(`${process.env.NEXT_PUBLIC_PB_PATH_API}/items/sub/category`);
        if (subCategory.status === 200) {
            subCategoryData = subCategory.data;
        }
    } catch (error) {
        subCategoryData = [];
    }



    return (
        <div className="flex-col">
            <div className="flex-1 space-y-4 p-8 pt-6">
                <ItemForm
                    initialData={itemData}
                    subCategory={subCategoryData}
                    initialSubCategory={initialSubCategoryData}
                />
            </div>
        </div>
    );
}

export default ItemPage;
