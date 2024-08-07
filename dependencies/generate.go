package dependencies

import (
	"fmt"
	"os"
	"strings"
)

func GenerateFolderStructure(folderName string) error {
	var initName string
	// Ambil kata terakhir dari folderName yang dipisahkan oleh "-"
	parts := strings.Split(folderName, "-")
	if len(parts) > 0 {
		initName = parts[len(parts)-1]
	}

	// 1. Buat folder utama
	if err := os.Mkdir(folderName, 0755); err != nil {
		return err
	}

	// 2. Definisikan struktur folder (sesuaikan dengan kebutuhan Anda)
	folderStructure := map[string]interface{}{
		initName: map[string]interface{}{
			"delivery": map[string]interface{}{
				"grpc": map[string]interface{}{
					"proto": map[string]interface{}{},
				},
				"http": map[string]interface{}{
					"handler": map[string]interface{}{
						"": []string{initName + ".go"},
					},
					"": []string{"api.go"},
				},
			},
			"repository": map[string]interface{}{
				"redis": map[string]interface{}{
					"": []string{initName + ".go"},
				},
				"sql": map[string]interface{}{},
			},
			"usecase": map[string]interface{}{
				"": []string{initName + ".go"},
			}},
		"app":    []string{"main.go"},
		"config": []string{"config.go"},
		"db":     map[string]interface{}{"migration": map[string]interface{}{}},
		"domain": []string{initName + ".go"},
		"helper": map[string]interface{}{},
		"":       []string{"openapi.yaml"},
	}

	// 3. Buat folder dan file secara rekursif
	if err := CreateStructure(folderName, folderStructure); err != nil {
		return err
	}

	// 4. Inisialisasi modul Go
	err := GoModInit(folderName)
	if err != nil {
		return err
	}
	fmt.Println("Modul Go berhasil diinisialisasi!")

	// 5. Tambahkan file config.go setelah struktur folder dibuat

	err = WriteConfigFile(folderName)
	if err != nil {
		return err
	}

	fmt.Println("Config file berhasil diinisalisasi!")

	err = WriteDomainFile(folderName, initName)
	if err != nil {
		return err
	}

	fmt.Println("Domain file berhasil diinisalisasi!")

	err = WriteDelivery(folderName, initName)
	if err != nil {
		return err
	}

	fmt.Println("Handler file berhasil diinisalisasi!")

	err = WriteRepository(folderName, initName)
	if err != nil {
		return err
	}

	fmt.Println("Repository file berhasil diinisalisasi!")

	err = WriteUseCase(folderName, initName)
	if err != nil {
		return err
	}

	fmt.Println("Usecase file berhasil diinisalisasi!")

	err = WriteMainFile(folderName, initName)
	if err != nil {
		return err
	}

	fmt.Println("Main file berhasil diinisalisasi!")

	// 6. Go mod tidy

	err = GoModTidy(folderName)
	if err != nil {
		return err
	}
	fmt.Println("Go mod tidy berhasil dijalankan")

	return nil
}
