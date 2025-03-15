package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/guddu75/goblog/internal/store"
)

var usernames = []string{
	"Aarav_99", "Diya_07", "Neha_Star", "Aryan_X", "Rohan_12",
	"Sanya_Pro", "Vikram_89", "Kabir_22", "Ananya_Rocks", "Meera_555",
	"Karthik_09", "Priya_Cool", "Aditya_77", "Simran_88", "Rahul_91",
	"Devansh_55", "Ishita_33", "Samar_007", "Tanvi_21", "Krishna_99",
	"Pooja_123", "Arjun_XY", "Sneha_456", "Yashraj_89", "Manish_777",
	"Vidya_999", "Ritika_101", "Nikhil_505", "Suhani_66", "Harshad_909",
	"Bhavya_55", "Vaani_07", "Sameer_321", "Tanya_008", "Uday_77",
	"Gautam_45", "Mansi_678", "Siddharth_95", "Radhika_010", "Veer_90",
	"Avni_89", "Kunal_300", "Jatin_65", "Ria_77", "Akash_33",
	"Tushar_007", "Sonal_000", "Dhruv_999", "Nandini_22", "Parth_87",
}

var titles = []string{
	"Understanding the Core Principles of Sanatan Dharma",
	"The Significance of Bhagavad Gita in Daily Life",
	"Upanishads: The Ultimate Knowledge of the Self",
	"Vedas: The Eternal Wisdom of Hinduism",
	"Bhakti Yoga: The Path of Devotion in Sanatan Dharma",
	"Karma and Reincarnation: Understanding the Cycle of Life",
	"Exploring the Four Yugas: Satya, Treta, Dwapara, and Kali",
	"Adi Shankaracharya and Advaita Vedanta Philosophy",
	"Why Do Hindus Perform Pujas? The Science Behind Rituals",
	"The Importance of Meditation in Hinduism",
	"Gayatri Mantra: The Most Powerful Vedic Chant",
	"Hanuman Chalisa: The Divine Power of Lord Hanuman",
	"The Role of Dharma in Leading a Righteous Life",
	"Understanding the Concept of Moksha in Hinduism",
	"Mahabharata: Lessons on Duty, Morality, and Dharma",
	"Ramayana: The Ideal Life of Lord Rama",
	"The Significance of OM in Hinduism",
	"Why Do Hindus Worship Different Gods? The Concept of Brahman",
	"Shiva Tandava Stotram: The Power and Energy of Lord Shiva",
	"The Science Behind Fasting in Sanatan Dharma",
	"Ekadashi Fasting: Spiritual and Scientific Benefits",
	"Sri Krishna’s Teachings on Dharma and Life",
	"Devi Mahatmya: The Divine Feminine in Hinduism",
	"The Significance of Rudraksha in Hindu Tradition",
	"The Importance of Guru in Sanatan Dharma",
	"Jyotirlingas: The Sacred Shrines of Lord Shiva",
	"Navaratri: The Nine Nights of Divine Energy",
	"Understanding Hindu Cosmology: The Creation and Dissolution",
	"The 108 Names of Lord Vishnu and Their Meaning",
	"Concept of Prarabdha and Sanchita Karma in Hinduism",
	"Yantras: The Sacred Geometries of Hinduism",
	"Why Do Hindus Do Namaste? The Spiritual Meaning",
	"Scientific Aspects of Temple Architecture in Hinduism",
	"Shri Ram Rajya: The Ideal Governance According to Hinduism",
	"The Role of Ayurveda in Hindu Dharma",
	"The Meaning of Trishul and Damaru in Lord Shiva’s Hands",
	"Power of Mantras: How Sound Vibration Affects Consciousness",
	"The Science of Aarti and Camphor in Hindu Worship",
	"Kali Yuga: The Present Age and Its Challenges",
	"Why Hindus Light a Lamp (Diya) During Worship?",
	"Understanding the 4 Purusharthas: Dharma, Artha, Kama, Moksha",
	"Sanatan Dharma and the Science of Rebirth",
	"The Concept of Ishta Devata in Hindu Worship",
	"The Role of Panchang (Hindu Calendar) in Daily Life",
	"Why Do Hindus Take a Dip in Holy Rivers?",
	"Significance of Chaturmas in Hindu Tradition",
	"The Role of Sanskrit in Preserving Vedic Knowledge",
	"The Concept of Kundalini Energy in Hinduism",
	"Why is Cow Sacred in Hinduism?",
	"How to Live According to Sanatan Dharma in Modern Life",
}

var contents = []string{
	"Sanatan Dharma, also known as Hinduism, is one of the world's oldest religions, emphasizing the eternal principles of righteousness, karma, and dharma.",
	"The Bhagavad Gita provides profound wisdom on duty, righteousness, and selfless action, guiding individuals toward liberation.",
	"The Upanishads explore the nature of the self (Atman) and its connection to the ultimate reality (Brahman).",
	"The Vedas, composed over 5,000 years ago, serve as the foundation of Hindu philosophy, rituals, and spiritual knowledge.",
	"Bhakti Yoga is the path of love and devotion, allowing devotees to connect with the divine through selfless surrender.",
	"Karma and reincarnation explain how actions in one life determine one's fate in future births, promoting moral responsibility.",
	"The four Yugas define the cosmic cycle of time, each with distinct characteristics affecting dharma and human behavior.",
	"Adi Shankaracharya's Advaita Vedanta philosophy emphasizes non-dualism, teaching that the individual soul (Atman) is one with Brahman.",
	"Hindu rituals and pujas are deeply rooted in science, promoting mental peace, focus, and spiritual growth.",
	"Meditation in Hinduism is a practice to attain self-realization, inner peace, and a connection with the divine.",
	"The Gayatri Mantra, considered the most powerful Vedic chant, invokes divine wisdom and removes ignorance.",
	"The Hanuman Chalisa is a devotional hymn dedicated to Lord Hanuman, known for its immense spiritual power and protection.",
	"Dharma is the righteous path that governs ethical and moral conduct in personal and societal life.",
	"Moksha is the ultimate goal of human existence, signifying liberation from the cycle of birth and death.",
	"The Mahabharata teaches valuable lessons on ethics, dharma, and the consequences of one's actions.",
	"The Ramayana narrates the life of Lord Rama, highlighting the ideals of righteousness, devotion, and duty.",
	"The sacred syllable 'OM' represents the essence of the universe and is used in meditation and prayers.",
	"In Hinduism, multiple gods are worshipped as different manifestations of the singular ultimate reality, Brahman.",
	"Shiva Tandava Stotram glorifies Lord Shiva’s cosmic dance, symbolizing destruction and renewal of the universe.",
	"Fasting in Hindu tradition has both spiritual and scientific benefits, aiding in physical purification and mental clarity.",
}

var tags = []string{
	"#SanatanDharma", "#Hinduism", "#BhagavadGita", "#Vedas", "#Upanishads",
	"#Karma", "#Dharma", "#Moksha", "#Yoga", "#Meditation", "#Spirituality",
	"#AdvaitaVedanta", "#HinduPhilosophy", "#Puja", "#HinduRituals", "#Mantras",
	"#HinduGods", "#AncientWisdom", "#IndianCulture", "#MahabharataLessons",
	"#RamayanaTeachings", "#VedicKnowledge", "#UpanishadicWisdom", "#AncientScriptures",
	"#NonDualism", "#Reincarnation", "#LawOfKarma", "#PujaVidhi", "#Aarti",
	"#HinduFasting", "#TempleTraditions", "#SacredMantras", "#BhaktiYoga",
	"#KarmaYoga", "#JnanaYoga", "#RajaYoga", "#HinduCosmology", "#Jyotirlinga",
	"#ShivaBhakti", "#KrishnaConsciousness", "#ShaktiWorship", "#OmChanting",
	"#HanumanBhakti", "#DivineEnergy", "#SacredChants", "#SanatanTraditions",
	"#HinduFestivals", "#SpiritualAwakening", "#SacredScriptures", "#IndianPhilosophy",
}

var comments = []string{
	"Great insights! 🙏",
	"Very informative post! 📖",
	"Jai Shri Ram! 🚩",
	"Well explained, thanks! 🙌",
	"Love this perspective! ❤️",
	"Such deep wisdom! 🕉️",
	"Amazing content! Keep it up! 👍",
	"This was really helpful! 😊",
	"Sanatan Dharma is eternal! 🔥",
	"Beautifully written! ✨",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)

	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, nil, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user: ", err)
			return
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post: ", post)
			return
		}
	}

	comments := generateComments(500, users, posts)

	for _, comment := range comments {
		if err := store.Commnets.Create(ctx, comment); err != nil {
			log.Println("Error creating comment: ", comment)
			return
		}
	}

	fmt.Println("Seeding complete")

}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", rand.Intn(1e5)),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Role:     store.Role{Name: "user"},
			//Password: "123123",
			// Password: &store.Password{
			// 	Text: &"123123",
			// },
		}
	}

	return users

}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {

		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Content: contents[rand.Intn(len(contents))],
			Title:   titles[rand.Intn(len(titles))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts

}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {

		post := posts[rand.Intn(len(posts))]
		user := users[rand.Intn(len(users))]

		cms[i] = &store.Comment{
			PostID:  post.ID,
			UserID:  user.ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
