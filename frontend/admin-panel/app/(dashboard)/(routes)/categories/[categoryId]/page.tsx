import { Category } from '@/app/schema/schema';
import { CategoryForm } from './components/category-form';
import axios from 'axios';

const CategoryPage = async ({ params }: { params: { categoryId: number } }) => {
    let categoryData: Category | null = null;


    if (params.categoryId) {
        try {
            const response = await axios.get(`${process.env.NEXT_PUBLIC_PB_PATH_API}/items/category/${params.categoryId}`);
            if (response.status === 200) {
                categoryData = response.data;
            }
        } catch (error) {
            console.error("Failed to fetch category:", error);
        }
    }

    return (
        <div className="flex-col">
            <div className="flex-1 space-y-4 p-8 pt-6">
                <CategoryForm initialData={categoryData} />
            </div>
        </div>
    );
}

export default CategoryPage;
