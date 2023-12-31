package locale

// FR locale validation message.
var FR = map[string]any{
	"required":        "Le champ %s est requis.",
	"string":          "Le champ %s doit être une chaîne de caractères.",
	"alpha":           "Le champ %s ne peut contenir que des lettres.",
	"numeric":         "Le champ %s doit être un nombre.",
	"alpha_numeric":   "Le champ %s ne peut contenir que des lettres et des chiffres.",
	"int":             "Le champ %s doit être un entier.",
	"uint":            "Le champ %s doit être un entier positif.",
	"float":           "Le champ %s doit être un nombre décimal.",
	"email":           "Le champ %s doit être une adresse email valide.",
	"phone":           "Le champ %s doit être un numéro de téléphone valide.",
	"phone_with_code": "Le champ %s doit être un numéro de téléphone valide avec le code du pays.",
	"username":        "Le champ %s doit être une adresse email valide, un numéro de téléphone valide ou un numéro de téléphone avec le code du pays.",
	"match":           "Le champ %s ne correspond pas.",
	"same":            "Les champs %s et %s doivent correspondre.",
	"unique":          "Le %s a déjà été pris.",
	"bool":            "Le champ %s doit être vrai.",
	"file":            "Le champ %s doit être un fichier.",
	"file_type":       "Le champ %s doit être un fichier de type : %s.",
	"image":           "Le champ %s doit être une image.",
	"image_type":      "Le champ %s doit être une image de type : %s.",
	"mimes":           "Le champ %s doit être un fichier de type : %s.",
	"gh_card":         "Le champ %s doit être une carte d'identité ghanéenne valide.",
	"gh_gps":          "Le champ %s doit être une adresse numérique ghanéenne valide.",
	"gt": map[string]string{
		"numeric": "Le champ %s doit être supérieur à %s.",
		"file":    "Le champ %s doit être supérieur à %s mégaoctets.",
		"string":  "Le champ %s doit comporter plus de %s caractères.",
		"slice":   "Le champ %s doit contenir plus de %s éléments.",
	},
	"gte": map[string]string{
		"numeric": "Le champ %s doit être supérieur ou égal à %s.",
		"file":    "Le champ %s doit être supérieur ou égal à %s mégaoctets.",
		"string":  "Le champ %s doit comporter au moins %s caractères.",
		"slice":   "Le champ %s doit contenir %s éléments ou plus.",
	},
	"lt": map[string]string{
		"numeric": "Le champ %s doit être inférieur à %s.",
		"file":    "Le champ %s doit être inférieur à %s mégaoctets.",
		"string":  "Le champ %s doit comporter moins de %s caractères.",
		"slice":   "Le champ %s doit contenir moins de %s éléments.",
	},
	"lte": map[string]string{
		"numeric": "Le champ %s doit être inférieur ou égal à %s.",
		"file":    "Le champ %s doit être inférieur ou égal à %s mégaoctets.",
		"string":  "Le champ %s doit comporter %s caractères ou moins.",
		"slice":   "Le champ %s ne doit pas contenir plus de %s éléments.",
	},
	"min": map[string]string{
		"numeric": "Le champ %s doit être d'au moins %s",
		"file":    "Le champ %s doit être d'au moins %s mégaoctets.",
		"string":  "Le champ %s doit comporter au moins %s caractères.",
		"slice":   "Le champ %s doit contenir au moins %s éléments.",
	},
	"max": map[string]string{
		"numeric": "Le champ %s ne doit pas être supérieur à %s.",
		"file":    "Le champ %s ne doit pas dépasser %s mégaoctets.",
		"string":  "Le champ %s ne doit pas comporter plus de %s caractères.",
		"slice":   "Le champ %s ne doit pas contenir plus de %s éléments.",
	},
	"equal": map[string]string{
		"numeric": "Le champ %s doit être égal à %s.",
		"file":    "Le champ %s doit être égal à %s mégaoctets.",
		"string":  "Le champ %s doit comporter exactement %s caractères.",
		"slice":   "Le champ %s doit comporter exactement %s éléments.",
	},
	"between": map[string]string{
		"numeric": "Le champ %s doit être compris entre %s et %s.",
		"file":    "Le champ %s doit être compris entre %s et %s mégaoctets.",
		"string":  "Le champ %s doit comporter entre %s et %s caractères.",
		"slice":   "Le champ %s doit comporter entre %s et %s éléments.",
	},
	"from": map[string]string{
		"numeric": "Le champ %s doit être compris entre %s et %s.",
		"file":    "Le champ %s doit être compris entre %s et %s mégaoctets.",
		"string":  "Le champ %s doit comporter entre %s et %s caractères.",
		"slice":   "Le champ %s doit comporter entre %s et %s éléments.",
	},
	"size": map[string]string{
		"numeric": "Le champ %s doit avoir une taille de %s.",
		"file_kb": "Le champ %s doit avoir une taille de %s kilooctets.",
		"file_mb": "Le champ %s doit avoir une taille de %s mégaoctets.",
		"file_gb": "Le champ %s doit avoir une taille de %s gigaoctets.",
		"string":  "Le champ %s doit comporter %s caractères.",
		"slice":   "Le champ %s doit contenir %s éléments.",
	},
}
