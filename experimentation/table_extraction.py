import pdfplumber
import pandas as pd
from pathlib import Path
from pprint import pprint

def extract_tables_pdfplumber(pdf_path, output_dir=None):
    """Extract tables using pdfplumber - better for complex layouts"""
    
    if output_dir is None:
        output_dir = Path(pdf_path).parent / "extracted_tables"
    
    output_dir = Path(output_dir)
    output_dir.mkdir(exist_ok=True)
    
    extracted_files = []
    table_count = 0
    
    header_re = r""

    with pdfplumber.open(pdf_path) as pdf:
        #pprint(pdf.pages[92].chars)
        #for char in pdf.pages[92].chars:
        #    if char['size'] > 10:
        #        print(char['text'],end="")
        #
        for page_num, page in enumerate(pdf.pages, 1):
            for char in page.chars:
                if char['size'] > 11:
                    print(char['text'],end="")
            print()

        #for page_num, page in enumerate(pdf.pages, 1):
            
            #tables = page.extract_tables()
            
            #for table in tables:
            #    if table:  # Check if table is not empty
            #        table_count += 1
            #        
            #        # Convert to DataFrame
            #        df = pd.DataFrame(table[1:], columns=table[0])  # First row as headers
            #        
            #        # Clean up
            #        df = df.dropna(how='all')
            #        
            #        output_file = output_dir / f"page_{page_num}_table_{table_count}.csv"
            #        df.to_csv(output_file, index=False)
            #        
            #        print(f"Page {page_num}, Table {table_count}: {len(df)} rows")
            #        extracted_files.append(output_file)
    
    return extracted_files
# Usage
if __name__ == "__main__":
    pdf_file = "DSP0134_3.3.0.pdf"  # Replace with your PDF path
    extract_tables_pdfplumber(pdf_file)
    #
    #print(f"\nExtracted {len(extracted)} tables successfully")
